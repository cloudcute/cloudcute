package app

import (
	"cloudcute/src/package/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func parseArgs() {
	flag.StringVar(&configPath, "c", config.GetDefaultConfigPath(), "配置文件路径")
	flag.Parse()
}

// Init 初始化
func Init() {
	go checkUpdate()
	parseArgs()
	config.Init(configPath)
	initApp()
	initStatic()
	if !config.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func initApp()  {
	fmt.Print(`
================================================
====                                        ====
====            Cloud Cute                  ====
====                                        ====
================================================
`)
}

func initStatic()  {
	// TODO 启动静态资源
}
