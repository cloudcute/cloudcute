package routers

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	return gin.Default()
}
