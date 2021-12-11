package app

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/utils"
	"cloudcute/src/routers"
	"github.com/gin-gonic/gin"
)

func Start() {
	var r = routers.Init()
	if config.SSLConfig.OpenSSL {
		startSSLListen(r)
	}else{
		startListen(r)
	}
}

func startListen(r *gin.Engine) {
	var listen = config.SystemConfig.Listen
	log.Info("开始监听 %s", listen)
	if err := r.Run(listen); err != nil {
		log.Error("监听错误[%s]，%s", listen, err)
		utils.WaitExit()
	}
}

func startSSLListen(r *gin.Engine) {
	var listen = config.SSLConfig.Listen
	var certPath = config.SSLConfig.CertPath
	var keyPath = config.SSLConfig.KeyPath
	if listen == "" || certPath == "" || keyPath == "" {
		log.Error("SSL配置为空")
		utils.WaitExit()
		return
	}
	go func() {
		log.Info("开始监听 %s", listen)
		if err := r.RunTLS(listen, certPath, keyPath); err != nil {
			log.Error("无法监听[%s]，%s", listen, err)
			utils.WaitExit()
		}
	}()
}
