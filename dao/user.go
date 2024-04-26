package dao

import (
	"aslon/global"
	"aslon/model/do"
)

type UserDao struct {
}

func (u *UserDao) GetUser(name, password string) (user do.User, err error) {
	sql := `SELECT id,name,password FROM user WHERE name=? AND password=?`
	err = global.UserDB.Raw(sql, name, password).Scan(&user).Error
	return
}

func (u *UserDao) CheckUser(name string) (user do.User, err error) {
	sql := `SELECT id,name FROM user WHERE name=?`
	err = global.UserDB.Raw(sql, name).Scan(&user).Error
	return
}

func (u *UserDao) InsertUser(id uint64, name, password, nick string) (err error) {
	user := do.User{Id: id, Name: name, Password: password, Nick: nick}
	err = global.UserDB.Table("user").Create(&user).Error
	return
}
