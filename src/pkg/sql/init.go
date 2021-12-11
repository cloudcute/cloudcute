package sql

import "cloudcute/src/pkg/log"

func initData() {
	if verifyVersion() {
		return
	}
	log.Info("初始化数据库数据...")

	log.Info("数据库初始化完成")
}

func verifyVersion() bool {
	//var setting Setting
	//return DB.Where("name = ?", "db_version_"+conf.RequiredDBVersion).First(&setting).Error != nil
	return false
}
