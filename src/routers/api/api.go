package api

import "github.com/gin-gonic/gin"

// UrlPrefix 访问api路径前缀
const UrlPrefix = "/api"

func Init(r *gin.Engine)  {
	var api = r.Group("/api")
	initApi(api)
}

func initApi(api *gin.RouterGroup) {

}
