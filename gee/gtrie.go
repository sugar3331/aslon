package gee

import "strings"

type node struct {
	pattern  string  //从跟节点到待匹配路由，例如/p/:lang
	part     string  //路由中的一部分，例如:lang
	children []*node //子节点，例如【doc，tutorial，intro】
	isWild   bool    // 是否精确匹配，part 含有：或*时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	//查看下一层,路由段有没有和节点段对应的或者是模糊匹配
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	// 遍历树的子节点，如果路由段和传入的节点段相同或者模糊匹配，返回匹配到的节点
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	//如果匹配到的树层高等于传入高度值，直接返回结束插入操作。
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]       //拿到路由节点
	child := n.matchChild(part) //定位到匹配到的子节点
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	//递归下一层插入
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	//退出条件 搜索高度到达目标树的层高
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	// 遍历匹配到的子节点进行递归操作，如果匹配到的子节点为空，直接退出
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
