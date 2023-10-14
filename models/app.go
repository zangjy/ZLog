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
	IsDel      int `gorm:"default:0"`
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

//
// DeleteApp
//  @Description: 删除APP
//  @param appId
//  @return bool
//
func DeleteApp(appId string) bool {
	tx := dao.DB.Table("app").Where("app_id = ?", appId).Update("is_del", 1)
	return tx.Error == nil
}

//
// GetAppList
//  @Description: 获取APP列表
//  @param page
//  @return int
//  @return []GetAppListInfoStruct
//
func GetAppList(page int) (int64, []GetAppListInfoStruct) {
	var result []GetAppListInfoStruct
	db := dao.DB.Table("app").Where("is_del = ?", 0)

	pageSize := 10
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	db.Find(&result)

	count := int64(0)
	db.Model(&GetAppListInfoStruct{}).Count(&count)

	return count, result
}
