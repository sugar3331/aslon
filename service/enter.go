package service

import "aslon/dao"

type ServiceGroup struct {
	UserService
	ChatService
}

var ServiceGroupApp = new(ServiceGroup)

var userDao = dao.DAOGroupApp.UserDao
