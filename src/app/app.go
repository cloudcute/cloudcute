package app

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/file_util"
	"cloudcute/src/routers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Start() {
	var r = routers.Init()
	//// 如果启用了SSL
	//if conf.SSLConfig.CertPath != "" {
	//	go func() {
	//		util.Log().Info("开始监听 %s", conf.SSLConfig.Listen)
	//		if err := api.RunTLS(conf.SSLConfig.Listen,
	//			conf.SSLConfig.CertPath, conf.SSLConfig.KeyPath); err != nil {
	//			util.Log().Error("无法监听[%s]，%s", conf.SSLConfig.Listen, err)
	//		}
	//	}()
	//}
	initStatic(r)
	var listen = config.SystemConfig.Listen
	log.Info("开始监听 %s", listen)
	if err := r.Run(listen); err != nil {
		log.Error("监听错误[%s]，%s", listen, err)
	}
}

func initStatic(r *gin.Engine)  {
	if !config.SystemConfig.OpenWeb {
		return
	}
	var staticPath string
	if config.IsDev {
		staticPath = "public/build"
	}else{
		staticPath = "public"
	}
	if !file_util.Exists(staticPath) {
		log.Error("未找到静态资源：%s", staticPath)
		return
	}
	r.Use(static.Serve("/", static.LocalFile(staticPath, false)))
}
