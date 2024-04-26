package v1

import "aslon/service"

type GroupApi struct {
	UserApi
	ChatApi
}

var GroupApiApp = new(GroupApi)
var (
	userService = service.ServiceGroupApp.UserService
	chatService = service.ServiceGroupApp.ChatService
)
