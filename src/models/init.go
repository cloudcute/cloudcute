package models

import (
	"cloudcute/src/models/setting"
	"cloudcute/src/models/user"
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/sql"
)

func Init() {
	sql.Init(getSqlConfig())
	initData()
}

func initData() {
	if verifyVersion() {
		return
	}
	log.Info("初始化数据...")
	startInitData()
	log.Info("数据初始化完成")
}

func verifyVersion() bool {
	if config.IsDev {
		// 开发模式经常改动, 每次都初始化
		return false
	}
	var s setting.Setting
	var err = sql.FirstQuery("name", setting.DBVersionKey, &s)
	return err == nil
}

func getSqlConfig() *sql.Config {
	return &sql.Config{
		Type:        config.DatabaseConfig.Type,
		Host:        config.DatabaseConfig.Host,
		User:        config.DatabaseConfig.User,
		Name:        config.DatabaseConfig.Name,
		Port:        config.DatabaseConfig.Port,
		DBPath:      config.DatabaseConfig.DBPath,
		Charset:     config.DatabaseConfig.Charset,
		Password:    config.DatabaseConfig.Password,
		TablePrefix: config.DatabaseConfig.TablePrefix,
	}
}

func startInitData() {
	reInitTable()
	setting.InitSettingData()
}

func reInitTable() {
	sql.ReInitTable(&setting.Setting{}, &user.User{}, &user.Group{})
}
