package service

import (
	"aslon/global"
	"aslon/middleware"
	"aslon/model/bo"
	"aslon/model/dto"
	"log"
	"strconv"
)

type UserService struct {
}

func (u *UserService) Login(user dto.LoginReq) (token bo.Token, err error) {
	dbRet, err := userDao.GetUser(user.Name, user.Password)
	if err != nil {
		log.Println(err)
		return
	}

	id, _ := strconv.ParseInt(strconv.FormatUint(dbRet.Id, 10), 10, 64)
	aToken, err := middleware.Token.GetToken(uint64(id), dbRet.Name, dbRet.Name, "user")

	if err != nil {
		log.Println(err)
		return
	}
	token.AccessToken = aToken
	return
}

func (u *UserService) Register(user dto.RegisterReq) (bo bool, err error) {
	dbRet, err := userDao.CheckUser(user.Name)
	if err != nil {
		log.Println(err)
		return
	}

	if dbRet.Id != 0 {
		return true, nil
	}

	// 生成ID
	id := uint64(global.SnowflakeNode.Generate())
	err = userDao.InsertUser(id, user.Name, user.Password, user.Nick)
	if err != nil {
		log.Println("插入用户数据失败", err)
	}
	return
}
