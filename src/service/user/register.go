package user

import (
	"cloudcute/src/models/define"
	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	UserName string `form:"username" json:"username" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=4,max=64"`
}

func (data *RegisterData) Register(c *gin.Context) define.Response {
	return define.GetResponse(nil, "ok")
}
