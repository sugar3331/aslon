package v1

import (
	"aslon/gee"
	"aslon/model/dto"
	"aslon/model/dto/response"
	"encoding/json"
	"net/http"
)

type UserApi struct {
}

func (u *UserApi) Login(c *gee.Context) {
	var user dto.LoginReq
	err := json.NewDecoder(c.Req.Body).Decode(&user)
	if err != nil {
		// 处理错误
		response.NormResponseWithBadInternal("bind data err", c)
		return
	}

	token, err := userService.Login(user)
	if err != nil {
		response.NormResponseWithBadInternal("登陆失败", c)
		return
	}
	response.NormResponseWithOK(token, "登陆成功", c)
}

func (u *UserApi) Register(c *gee.Context) {
	var user dto.RegisterReq
	err := json.NewDecoder(c.Req.Body).Decode(&user)
	if err != nil {
		// 处理错误
		response.NormResponseWithBadInternal("bind data err", c)
		return
	}

	in, err := userService.Register(user)
	if err != nil {
		response.NormResponseWithBadInternal("注册失败", c)
		return
	}
	if in {
		response.NormResponseWithBadInternal("注册失败,用户名已存在请重新输入", c)
		return
	}
	response.NormResponseWithOK("", "注册成功", c)
}

func (u *UserApi) GetLogin(c *gee.Context) {
	c.HTML(http.StatusOK, "login.html", gee.H{})
}

func (u *UserApi) GetRegister(c *gee.Context) {
	c.HTML(http.StatusOK, "register.html", gee.H{})
}
