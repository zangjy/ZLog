package models

import (
	"ZLog/dao"
)

type User struct {
	Id       uint   `gorm:"primaryKey"`
	UserName string `gorm:"unique"`
	Password string
}

//
// GetUserInfo
//  @Description: 根据用户名和密码获取用户信息
//  @param userName
//  @param pwd
//  @return error
//  @return *User
//
func GetUserInfo(userName, pwd string) (error, *User) {
	user := &User{}
	tx := dao.DB.Table("user").Where("user_name = ? and password = ?", userName, pwd).First(user)
	return tx.Error, user
}
