package app

import (
	"cloudcute/src/package/config"
	"cloudcute/src/package/log"
	"cloudcute/src/routers"
)

func Start() {
	var api = routers.Init()
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
	var listen = config.SystemConfig.Listen
	log.Info("开始监听 %s", listen)
	if err := api.Run(listen); err != nil {
		log.Error("监听错误[%s]，%s", listen, err)
	}
}
