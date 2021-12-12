package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	NickName  string `gorm:"size:100"`
	Password  string
	GroupId   uint
}
