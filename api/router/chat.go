package router

import (
	v1 "aslon/api/handlers/v1"
	"aslon/gee"
)

type ChatRouter struct{}

func (ch *ChatRouter) InitChatRouter(Router *gee.RouterGroup) {
	chatRouter := Router.Group("")
	chatApi := v1.GroupApiApp.ChatApi

	{
		chatRouter.GET("chat", chatApi.GetChat)

		chatRouter.GET("p2pChat", chatApi.GetP2PChat)

		chatRouter.GET("/", chatApi.GetPublicChat)
	}
}
