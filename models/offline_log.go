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

//
// GetTaskLog
//  @Description: 查询任务日志
//  @param input
//  @return int64
//  @return []GetTaskLogInfoStruct
//
func GetTaskLog(input GetTaskLogInputStruct) (int64, []GetTaskLogInfoStruct) {
	db := dao.DB.Table("offline_log").Where("1 = 1")

	if len(input.TaskId) > 0 {
		db = db.Where("task_id = ?", input.TaskId)
	}
	if len(input.SystemVersion) > 0 {
		db = db.Where("system_version LIKE ?", "%"+input.SystemVersion+"%")
	}
	if len(input.AppVersion) > 0 {
		db = db.Where("app_version LIKE ?", "%"+input.AppVersion+"%")
	}
	if input.StartStamp > 0 {
		db = db.Where("time_stamp >= ?", input.StartStamp)
	}
	if input.EndStamp > 0 {
		db = db.Where("time_stamp <= ?", input.EndStamp)
	}
	if input.LogLevel >= 0 {
		db = db.Where("log_level = ?", input.LogLevel)
	}
	if len(input.Identify) > 0 {
		db = db.Where("identify LIKE ?", "%"+input.Identify+"%")
	}
	if len(input.Tag) > 0 {
		db = db.Where("tag LIKE ?", "%"+input.Tag+"%")
	}
	if len(input.Msg) > 0 {
		db = db.Where("msg LIKE ?", "%"+input.Msg+"%")
	}

	db = db.Order("time_stamp DESC")

	pageSize := 10
	if input.Page < 1 {
		input.Page = 1
	}
	offset := (input.Page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	var count int64
	db.Model(&GetTaskLogInfoStruct{}).Count(&count)

	var logs []GetTaskLogInfoStruct
	db.Select("sequence, system_version, app_version, time_stamp, log_level, identify, tag, msg").Find(&logs)

	if len(logs) == 0 {
		logs = make([]GetTaskLogInfoStruct, 0)
	}

	return count, logs
}
