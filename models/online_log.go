package models

import "ZLog/dao"

type OnlineLog struct {
	Id            uint `gorm:"primaryKey"`
	SessionId     string
	Sequence      int64
	SystemVersion string
	AppVersion    string
	TimeStamp     int64
	LogLevel      int `gorm:"comment:0:INFO 1:DEBUG 2:VERBOSE 3:WARN 4:ERROR"`
	Identify      string
	Tag           string
	Msg           string
}

//
// WriteOnlineLogs
//  @Description: 批量写入数据
//  @param logs
//  @return error
//
func WriteOnlineLogs(logs []*OnlineLog) error {
	//批量写入数据
	if err := dao.DB.Table("online_log").Create(logs).Error; err != nil {
		return err
	}
	return nil
}

//
// GetDeviceLogs
//  @Description: 查询设备日志
//  @param input
//  @return []GetDeviceLogInfoStruct
//
func GetDeviceLogs(input GetDeviceLogInputStruct) []GetDeviceLogInfoStruct {
	db := dao.DB.Table("online_log").Where("1 = 1")

	if len(input.SessionId) > 0 {
		db = db.Where("session_id = ?", input.SessionId)
	}
	if len(input.SystemVersion) > 0 {
		db = db.Where("system_version LIKE ?", "%"+input.SystemVersion+"%")
	}
	if len(input.AppVersion) > 0 {
		db = db.Where("app_version LIKE ?", "%"+input.AppVersion+"%")
	}
	if input.StartStamp != 0 {
		db = db.Where("time_stamp >= ?", input.StartStamp)
	}
	if input.EndStamp != 0 {
		db = db.Where("time_stamp <= ?", input.EndStamp)
	}
	if input.LogLevel != 0 {
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

	pageSize := 10
	if input.Page < 1 {
		input.Page = 1
	}
	offset := (input.Page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	var logs []GetDeviceLogInfoStruct
	db.Select("sequence, system_version, app_version, time_stamp, log_level, identify, tag, msg").Find(&logs)

	if len(logs) == 0 {
		logs = make([]GetDeviceLogInfoStruct, 0)
	}

	return logs
}
