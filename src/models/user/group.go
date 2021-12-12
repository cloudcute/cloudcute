package user

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name          string
}

const GroupSuperAdmin uint = 1
const GroupAdmin uint = 2
const GroupUser uint = 3

var GroupNames = map[uint] string {
	GroupSuperAdmin : "超级管理员",
	GroupAdmin : "管理员",
	GroupUser : "用户",
}
