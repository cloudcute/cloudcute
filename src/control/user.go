package control

import (
	"cloudcute/src/service/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data user.LoginData
	startHandleJSON(c, &data, data.Login)
}

func Register(c *gin.Context) {
	var data user.RegisterData
	startHandleJSON(c, &data, data.Register)
}

func Logout(c *gin.Context) {
	startHandle(c, user.Logout)
}
