package sql

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/path_util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"xorm.io/xorm"
	xormLog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var sqlConfig *Config
var engine *xorm.Engine

func Init(c *Config) {
	log.Info("连接数据库...")
	sqlConfig = c
	var t = sqlConfig.Type
	var user = sqlConfig.User
	var pwd = sqlConfig.Password
	var host = sqlConfig.Host
	var port = sqlConfig.Port
	var name = sqlConfig.Name
	var charset = sqlConfig.Charset
	var dbPath = sqlConfig.DBPath
	var prefix = sqlConfig.TablePrefix
	var err error
	switch t {
	case "", "sqlite", "sqlite3":
		engine, err = xorm.NewEngine("sqlite3", path_util.GetAbsPath(dbPath))
	case "mysql":
		// "root:123@/test?charset=utf8&parseTime=True&loc=Local"
		engine, err = xorm.NewEngine(t, fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			user, pwd, host, port, name, charset))
	case "postgres":
		engine, err = xorm.NewEngine(t, fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			host, user, pwd, name, port))
	default:
		log.Panic("不支持的数据库类型: %s", t)
	}
	if err != nil {
		log.Panic("数据库连接失败: %s", err)
	}
	// 处理前缀
	engine.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, prefix))
	// 设置log
	setLog()
	// 设置连接池
	engine.DB().SetMaxIdleConns(50)
	engine.DB().SetMaxOpenConns(100)
	engine.DB().SetConnMaxLifetime(time.Second * 30)

	// 初始化数据
	initData()
}

func setLog() {
	if config.SystemConfig.Debug {
		engine.SetLogLevel(xormLog.LOG_DEBUG)
	}else{
		engine.SetLogLevel(xormLog.LOG_ERR)
	}
}
