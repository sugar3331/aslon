package initialize

import (
	"aslon/api/router"
	"aslon/gee"
	"aslon/global"
	"aslon/middleware"
	"log"
)

// 初始化总路由

func Routers() *gee.Engine {

	Router := gee.New()
	Router.Use(gee.Recovery())
	Router.LoadHTMLGlob("static/*")

	userRouter := router.RouterGroupApp.UserRouter
	chatRouter := router.RouterGroupApp.ChatRouter
	// 如果想要不使用ngeex代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/static")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	//Router.StaticFS(global.BuddyConfig.Local.StorePath, http.Dir(global.BuddyConfig.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/rpc.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	log.Printf("use middleware cors")
	//docs.SwaggerInfo.BasePath = global.BuddyConfig.System.RouterPrefix
	//Router.GET(global.BuddyConfig.System.RouterPrefix+"/swagger/*any", geeSwagger.WrapHandler(swaggerFiles.Handler))
	//global.BuddyLog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GlobalConfig.System.RouterPrefix)
	//{
	//	// 健康监测
	//	PublicGroup.GET("/health", func(c *gee.Context) {
	//		c.JSON(http.StatusOK, "ok")
	//	})
	//}
	{
		userRouter.InitUserRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.GlobalConfig.System.RouterPrefix)
	PrivateGroup.Use()
	{
		chatRouter.InitChatRouter(PublicGroup)
	}

	//global.BuddyLog.Info("router register success")
	return Router
}
