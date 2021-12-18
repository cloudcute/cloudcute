package user

import (
	"cloudcute/src/models/define"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	UserName string `form:"username" json:"username" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=4,max=64"`
}

func (data *LoginData) Login(c *gin.Context) define.Response {
	return define.GetResponse(nil, "ok")
}


