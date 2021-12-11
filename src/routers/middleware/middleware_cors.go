package middleware

import (
	"cloudcute/src/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initCORS(r *gin.Engine) {
	if !config.CORSConfig.OpenCORS {
		return
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORSConfig.AllowOrigins,
		AllowMethods:     config.CORSConfig.AllowMethods,
		AllowHeaders:     config.CORSConfig.AllowHeaders,
		AllowCredentials: config.CORSConfig.AllowCredentials,
		ExposeHeaders:    config.CORSConfig.ExposeHeaders,
	}))
}
