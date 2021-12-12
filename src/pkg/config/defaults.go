package config

const AppName = "cloudcute"

// SystemConfig 系统公用配置
var SystemConfig = &system{
	Debug:  false,
	Listen: ":5666",
	OpenWeb: true,
}

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:        "",
	Charset:     "utf8",
	DBPath:      AppName + ".db",
	Port:        3306,
	TablePrefix: AppName + "_",
	Name:        AppName,
	Host:        "localhost",
}

// SSLConfig SSL配置
var SSLConfig = &ssl{
	OpenSSL:  false,
	CertPath: "",
	KeyPath:  "",
	Listen:   ":443",
}

// CORSConfig 跨域配置
var CORSConfig = &cors{
	OpenCORS:         false,
	AllowOrigins:     []string{""},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
	AllowHeaders:     []string{"Cookie", "X-Cr-Policy", "Authorization", "Content-Length", "Content-Type", "X-Cr-Path", "X-Cr-FileName"},
	AllowCredentials: false,
	ExposeHeaders:    nil,
}
