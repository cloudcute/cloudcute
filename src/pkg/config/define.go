package config

// system 系统配置
type system struct {
	Listen        string `validate:"required"`
	Debug         bool
	OpenWeb       bool
}

// database 数据库
type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBPath      string
	Port        int
	Charset     string
}

// ssl ssl配置
type ssl struct {
	OpenSSL  bool
	CertPath string `validate:"omitempty,required"`
	KeyPath  string `validate:"omitempty,required"`
	Listen   string `validate:"required"`
}

// 跨域配置
type cors struct {
	OpenCORS         bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
}
