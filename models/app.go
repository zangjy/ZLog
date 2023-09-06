package models

import (
	"ZLog/dao"
	"ZLog/utils"
	"time"
)

type App struct {
	Id         uint `gorm:"primaryKey"`
	AppName    string
	AppId      string `gorm:"unique"`
	CreateTime int64
}

//
// CreateApp
//  @Description: 创建APP
//  @param appName
//  @return error
//  @return string
//
func CreateApp(appName string) (error, string) {
	appId := utils.WorkerInstance.GetId()
	tx := dao.DB.Table("app").Create(&App{AppName: appName, AppId: appId, CreateTime: time.Now().Unix()})
	return tx.Error, appId
}
