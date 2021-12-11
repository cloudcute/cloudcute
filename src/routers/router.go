package routers

import (
	"cloudcute/src/routers/api"
	"cloudcute/src/routers/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	var r = gin.Default()
	middleware.Init(r)
	api.InitApi(r.Group("/api"))
	return r
}
