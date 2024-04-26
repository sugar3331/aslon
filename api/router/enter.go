package router

type RouterGroup struct {
	ChatRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)
