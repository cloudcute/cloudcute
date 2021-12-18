package auth

import (
	"cloudcute/src/models/define"
	"cloudcute/src/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Token struct {
	Token *token.Token
	Err error
}

func UserAuthMiddleInit(r *gin.RouterGroup) {
	r.Use(getTokenMiddleware())
	r.Use(authUser())
}

func authUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t = c.MustGet(define.TokenKey)
		var token = t.(*Token)
		if token.Err != nil {
			var output = define.CreateErrorResponse(define.ErrorNotLogin, "解析token失败", token.Err.Error())
			c.JSON(http.StatusOK, output)
			c.Abort()
			return
		}
		c.Set(define.UserInfoKey, token.Token) // 顺利得到token中的用户信息
	}
}

// GetTokenMiddleware 获取Token中间件
func getTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString = c.Request.Header.Get(define.TokenHeaderKey)
		var token, err = token.ParseToken(tokenString)
		var t = Token{}
		if err != nil {
			t.Err = err
		}else{
			t.Token = token
		}
		c.Set(define.TokenKey, &t)
	}
}
