package app

import (
	"cloudcute/src/module/config"
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
   _____ _                 _    _____      _       
  / ____| |               | |  / ____|    | |      
 | |    | | ___  _   _  __| | | |    _   _| |_ ___ 
 | |    | |/ _ \| | | |/ _  | | |   | | | | __/ _ \
 | |____| | (_) | |_| | (_| | | |___| |_| | ||  __/
  \_____|_|\___/ \__,_|\__,_|  \_____\__,_|\__\___|



====================================================

`)
}

func initStatic()  {
	// TODO 启动静态资源
}
