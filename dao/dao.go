package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

//
// InitDao
//  @Description: 初始化数据库连接
//  @param userName 用户名
//  @param passWord 密码
//  @param host 主机名
//  @param port 端口号
//  @param dbName 数据库
//  @return err err不为nil是代表数据库连接失败
//
func InitDao(userName, passWord, host, port, dbName string) (err error) {
	dsn := userName + ":" + passWord + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	return err
}
