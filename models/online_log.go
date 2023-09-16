package models

import "ZLog/dao"

type OnlineLog struct {
	Id            uint `gorm:"primaryKey"`
	SessionId     string
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
