package user

import (
	"cloudcute/src/models/define"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) define.Response {
	return define.GetResponse(nil)
}
