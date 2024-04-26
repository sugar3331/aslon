package response

import (
	"aslon/gee"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Detail string      `json:"detail"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

type NResponse struct {
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Detail string      `json:"detail"`
}

const (
	NormSUCCESS             = "200"            //SUCCESS 成功
	NormUNAUTHORIZED        = "401"            //UNAUTHORIZED 未授权 ，请登录获取授权
	NormFORBIDDEN           = "403"            //FORBIDDEN 禁止的操作
	NormInternalServerError = "500"            //InternalServerError 内部服务错误
	NormBADPARM             = "BAD_PARAMETER"  //参数错误
	NormTOAST               = "TOAST"          //弹吐司，取 msg 内容进行展示
	NormRISK                = "RISK_TRIGGERED" //内容触发风控
	NormTOKENEXP            = "TOKEN_EXPIRED"  //TOKEN 过期，需要刷新 TOKEN
)

func NormResult(code string, data interface{}, msg string, detail string, c *gee.Context) {
	c.JSON(http.StatusOK, NResponse{
		Code:   code,
		Data:   data,
		Msg:    msg,
		Detail: detail,
	})
}

func NormResponse(code string, data interface{}, msg string, detail string, c *gee.Context) {
	NormResult(code, data, msg, detail, c)
}
func NormResponseWithOK(data interface{}, msg string, c *gee.Context) {
	NormResult(NormSUCCESS, data, msg, "", c)
}

func NormResponseWithBadInternal(msg string, c *gee.Context) {
	NormResult(NormInternalServerError, "", msg, "", c)
}

func NormResponseWithBadToken(msg string, c *gee.Context) {
	NormResult(NormUNAUTHORIZED, "", msg, "", c)
}

func NormResponseWithTokenExp(msg string, c *gee.Context) {
	NormResult(NormTOKENEXP, "", msg, "", c)
}

func NormResponseWithBadParam(msg string, c *gee.Context) {
	NormResult(NormBADPARM, nil, msg, "", c)
}
