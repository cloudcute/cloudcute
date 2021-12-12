package setting

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/sql"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Type  string `gorm:"not null"`
	Name  string `gorm:"not null;uniqueIndex:setting_key"`
	Value string `gorm:"size:65535"`
}

const DBVersionKey = "db_version" + config.DBVersion
const ServerVersionKey = "app_version"
const WebVersionKey = "web_version"

func InitSettingData() {
	var s = []Setting{
		{Name: DBVersionKey, Value: config.DBVersion, Type: "version"},
		{Name: ServerVersionKey, Value: config.ServerVersion, Type: "version"},
		{Name: WebVersionKey, Value: config.WEBVersion, Type: "version"},
	}
	for _, value := range s {
		_ = sql.Set("name", value.Name, &value)
	}
}
