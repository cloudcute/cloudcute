package user

import (
	"cloudcute/src/models/define"
	"cloudcute/src/models/user"
	"cloudcute/src/pkg/sql"
	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	UserName string `form:"username" json:"username" binding:"required,min=2,max=64"`
	Email    string `form:"email"    json:"email"    binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=4,max=64"`
}

func (data *RegisterData) Register(c *gin.Context) define.Response {
	var err error
	_, err = user.GetUserByUserName(data.UserName)
	if err == nil {
		return define.GetErrorResponseByStr("注册失败", "用户名已存在")
	}
	_, err = user.GetUserByEmail(data.Email)
	if err == nil {
		return define.GetErrorResponseByStr("注册失败", "邮箱已存在")
	}
	var userData = user.CreateUser(data.UserName, data.Email, data.Password)
	err = sql.Create(&userData)
	if err != nil {
		return define.GetErrorResponseByError("注册失败", err)
	}
	var result = make(map[string]interface{})
	userData.Password = ""
	result["user"] = userData
	// 注册成功
	return define.GetResponse(result)
}
