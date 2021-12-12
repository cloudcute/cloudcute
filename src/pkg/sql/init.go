package sql

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/path"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var sqlConfig *Config
var DB *gorm.DB

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
	var gormConfig = getGormConfig(prefix)
	var db *gorm.DB
	var err error
	switch t {
	case "", "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open(path.GetAbsPath(dbPath)), gormConfig)
	case "mysql":
		// "root:123@/test?charset=utf8&parseTime=True&loc=Local"
		var dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			user, pwd, host, port, name, charset)
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	case "postgres":
		var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			host, user, pwd, name, port)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	case "sqlserver":
		var dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			user, pwd, host, port, name)
		db, err = gorm.Open(sqlserver.Open(dsn), gormConfig)
	default:
		log.Panic("不支持的数据库类型: %s", t)
	}
	if err != nil {
		log.Panic("数据库连接失败: %s", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}

func getGormConfig(prefix string) *gorm.Config {
	var logMode logger.LogLevel
	if config.SystemConfig.Debug {
		logMode = logger.Error
	}else{
		logMode = logger.Silent
		// logMode = logger.Error
	}
	return &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,   // 表名前缀
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			// NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	}
}
