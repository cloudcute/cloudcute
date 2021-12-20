package user

import (
	"cloudcute/src/models/define"
	"cloudcute/src/models/user"
	"cloudcute/src/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

const(
	MethodUserName  = 0
	MethodEmail     = 1
)

type LoginData struct {
	Method   int    `form:"method" json:"method"`
	UserName string `form:"username" json:"username" binding:"min=2,max=64"`
	Email    string `form:"email"    json:"email"    binding:"min=4,email"`
	Password string `form:"password" json:"password" binding:"required,min=4,max=64"`
}

func (data *LoginData) Login(c *gin.Context) define.Response {
	var userData user.User
	var err error
	if data.Method == MethodUserName {
		userData, err = user.GetUserByUserName(data.UserName)
	}else if data.Method == MethodEmail {
		userData, err = user.GetUserByEmail(data.Email)
	}else{
		var errStr = fmt.Sprintf("不存在 method (%d)", data.Method)
		return define.GetErrorResponseByStr("登录失败", errStr)
	}
	if err != nil {
		var errStr = "用户名或密码错误"
		if config.IsDev {
			errStr = errStr + " : " + err.Error()
		}
		return define.GetErrorResponseByStr("登录失败", errStr)
	}
	if !userData.CheckPassword(data.Password) {
		return define.GetErrorResponseByStr("登录失败", "用户名或密码错误")
	}
	token, err := CreateToken(userData)
	if err != nil {
		return define.GetErrorResponseByStr("登录失败", "创建token失败: " + err.Error())
	}
	var result = make(map[string]interface{})
	userData.Password = ""
	result["user"] = userData
	result["token"] = token
	// 登录成功
	return define.GetResponse(result)
}
