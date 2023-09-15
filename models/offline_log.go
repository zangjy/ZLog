package models

import (
	"ZLog/dao"
)

type OfflineLog struct {
	Id            uint `gorm:"primaryKey"`
	TaskId        string
	Sequence      int64
	SystemVersion string
	AppVersion    string
	TimeStamp     int64
	LogLevel      int
	Identify      string
	Tag           string
	Msg           string
}

//
// WriteOfflineLogs
//  @Description: 批量写入数据
//  @param taskId
//  @param logs
//  @return error
//
func WriteOfflineLogs(logs []*OfflineLog) error {
	//批量写入数据
	if err := dao.DB.Table("offline_log").Create(logs).Error; err != nil {
		return err
	}
	return nil
}
