package router

import (
	v1 "aslon/api/handlers/v1"
	"aslon/gee"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gee.RouterGroup) {
	userRouter := Router.Group("")
	userApi := v1.GroupApiApp.UserApi
	{
		userRouter.GET("login", userApi.GetLogin)
		userRouter.POST("login", userApi.Login)

		userRouter.GET("register", userApi.GetRegister)
		userRouter.POST("register", userApi.Register)
	}

}
