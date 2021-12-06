package config

// system 系统配置
type system struct {
	Listen        string `validate:"required"`
	Debug         bool
}

// database 数据库
type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBFile      string
	Port        int
	Charset     string
}