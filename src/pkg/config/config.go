package config

import (
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/file_util"
	"cloudcute/src/pkg/utils/path_util"
	"github.com/go-playground/validator/v10"
	"gopkg.in/ini.v1"
)

var config *ini.File

const devMode = "develop"
var RuntimeMode = devMode
var IsDev = true

const defaultConfigName = "config.ini"
const defaultConfig = `[System]
Listen = :5666
OpenWeb = true
Debug = false
`

// Init 初始化配置文件
func Init(path string) {
	var err error
	if path == "" || !file_util.Exists(path) {
		if path == "" {
			path = GetDefaultConfigPath()
		}
		var confContent = defaultConfig
		var f, err = file_util.CreatNestedFile(path)
		if err != nil {
			log.Panic("无法创建配置文件, %s", err)
		}
		_, err = f.WriteString(confContent)
		if err != nil {
			log.Panic("无法写入配置文件, %s", err)
		}
		f.Close()
	}
	config, err = ini.Load(path)
	if err != nil {
		log.Panic("无法解析配置文件 '%s': %s", path, err)
	}
	var sections = map[string]interface{}{
		"Database":   DatabaseConfig,
		"System":     SystemConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			log.Panic("配置文件 %s 分区解析失败: %s", sectionName, err)
		}
	}

	// 重设log等级
	if !SystemConfig.Debug {
		log.SetLevel(log.LevelInformational)
	}

	IsDev = RuntimeMode == devMode
}

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := config.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}
	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}
	return nil
}

func GetDefaultConfigPath() string {
	return path_util.GetAbsPath(defaultConfigName)
}
