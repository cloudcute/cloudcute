package models

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/sql"
)

func Init() {
	sql.Init(getSqlConfig())
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
