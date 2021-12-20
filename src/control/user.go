package control

import (
	"cloudcute/src/service/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data user.LoginData
	var err = parseJSONData(c, &data)
	if err != nil {
		result(c, getParameRepByError(err))
		return
	}
	if data.Method == user.MethodUserName && data.UserName == "" {
		result(c, getParameRepByStr("username is null"))
		return
	}
	if data.Method == user.MethodEmail && data.Email == "" {
		result(c, getParameRepByStr("email is null"))
		return
	}
	result(c, data.Login(c))
}

func Register(c *gin.Context) {
	var data user.RegisterData
	if err := parseJSONData(c, &data); err != nil {
		result(c, getParameRepByError(err))
		return
	}
	result(c, data.Register(c))
}

func Logout(c *gin.Context) {
	result(c, user.Logout)
}
