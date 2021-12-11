package config

// SystemConfig 系统公用配置
var SystemConfig = &system{
	Debug:  false,
	Listen: ":5666",
	OpenWeb: true,
}

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:    "UNSET",
	Charset: "utf8",
	DBFile:  "cloudreve.db",
	Port:    3306,
}
