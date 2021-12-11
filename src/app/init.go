package app

import (
	"cloudcute/src/models"
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/update"
	"cloudcute/src/pkg/utils/path_util"
	"cloudcute/src/pkg/utils/utils"
	"cloudcute/src/routers/middleware"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	configPath string
	exportStatic bool
)

func parseArgs() {
	flag.StringVar(&configPath, "c", config.GetDefaultConfigPath(), "配置文件路径")
	flag.BoolVar(&exportStatic, "export", false, "导出静态资源")
	flag.Parse()
}

// Init 初始化
func Init() {
	parseArgs()
	initConfig()
	initLog()

	initCommand()

	update.CheckUpdate()
	initAppInfo()
	initMode()
	models.Init()
}

func initConfig()  {
	config.Init(configPath)
}

func initAppInfo()  {
	fmt.Print(`
   _____ _                 _    _____      _       
  / ____| |               | |  / ____|    | |      
 | |    | | ___  _   _  __| | | |    _   _| |_ ___ 
 | |    | |/ _ \| | | |/ _  | | |   | | | | __/ _ \
 | |____| | (_) | |_| | (_| | | |___| |_| | ||  __/
  \_____|_|\___/ \__,_|\__,_|  \_____\__,_|\__\___|



====================================================

`)
	if config.IsDev {
		log.Info("Dev Mode")
	}
}

func initMode()  {
	if !config.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func initLog()  {
	if !config.SystemConfig.Debug {
		log.SetLevel(log.LevelInfo)
	}
	gin.DisableConsoleColor()
	var logPath = path_util.GetAbsPath("gin_log.txt")
	var f, _ = os.Create(logPath)
	if !config.SystemConfig.Debug {
		gin.DefaultWriter = io.MultiWriter(f)
	}else {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}

func initCommand() {
	if exportStatic {
		middleware.ExportStatic()
		utils.WaitExit()
	}
}
