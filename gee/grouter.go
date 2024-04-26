package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

// 构建添加路由哈希
func (r *router) addRouter(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	//每个请求方法构建一颗树
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	//执行插入
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRouter(method, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)

	params := make(map[string]string)
	root, ok := r.roots[method]

	//如果找不到开始的树，直接退出
	if !ok {
		return nil, nil
	}

	//拿到根节点
	n := root.search(searchParts, 0)

	//从方法树开始查找构建
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 page not found:%s\n", c.Path)
		})
	}
	c.Next()
}
