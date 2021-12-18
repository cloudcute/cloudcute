package api

import (
	"cloudcute/src/control"
	"cloudcute/src/models/define"
	apiAuth "cloudcute/src/routers/api/auth"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine)  {
	var api = r.Group(define.APIPrefix)
	initApi(api)
}

func initApi(api *gin.RouterGroup) {

	var user = api.Group(define.APIUserPrefix)
	{
		user.POST("login", control.Login)
		user.POST("register", control.Register)
	}

	var auth = api.Group("") // 需要认证的api
	apiAuth.UserAuthMiddleInit(auth)
	{
		var user = auth.Group(define.APIUserPrefix)
		{
			user.POST("logout", control.Logout)
		}
	}
}
