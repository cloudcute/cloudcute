package user

import (
	"cloudcute/src/pkg/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(100);uniqueIndex"`
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	Password  string
	GroupId   uint
}

// GetUserByID 用ID获取用户
func GetUserByID(ID interface{}) (User, error) {
	var user User
	var err = sql.FirstPreload(ID, &user)
	return user, err
}

// GetUserByUserName 用UserName获取用户
func GetUserByUserName(username string) (User, error) {
	var user User
	var err = sql.FirstQueryPreload("user_name", username, &user)
	return user, err
}

// GetUserByEmail 用Email获取用户
func GetUserByEmail(email string) (User, error) {
	var user User
	var err = sql.FirstQueryPreload("email", email, &user)
	return user, err
}

// CreateUser 创建一个 User
func CreateUser(username string, email string, password string) User {
	return User{
		UserName: username,
		Email: email,
		Password: password,
		GroupId: GroupUser,
	}
}

// CheckPassword 根据明文校验密码
func (user *User) CheckPassword(password string) bool {
	return user.Password == password
}
