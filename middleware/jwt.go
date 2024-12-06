package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/consts"
	"github.com/kasiforce/trade/pkg/e"
	"github.com/kasiforce/trade/pkg/util"
	"strings"
)

// AuthToken 对前端发来的http请求验证token是否存在、正确、已过期等
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.TokenExpired
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "token不能为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(token, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			code = e.TokenExpired
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "token格式错误",
			})
			c.Abort()
			return
		}
		err := util.ParseToken(c, parts[1])
		if err != nil {
			code = e.TokenExpired
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": err.Error(),
			})
			c.Abort()
			return
		}
		//SetToken(c, newToken)
		c.Next()
	}
}

// SetToken 将token添加到http响应的Header
func SetToken(c *gin.Context, accessToken string) {
	secure := IsHttps(c)
	c.Header(consts.AccessTokenHeader, accessToken)
	c.SetCookie(consts.AccessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
}

func IsHttps(c *gin.Context) bool {
	if c.GetHeader(consts.HeaderForwardedProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
