package middleware

import (
	"aslon/gee"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Cors() gee.HandlerFunc {
	return func(c *gee.Context) {
		//method := c.Req.Method
		origin := c.Req.Header.Get("Origin")
		c.SetHeader("Access-Control-Allow-Origin", origin)
		c.SetHeader("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.SetHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.SetHeader("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.SetHeader("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}
		// 处理请求
		c.Next()
	}
}
