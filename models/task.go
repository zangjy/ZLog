package models

import (
	"ZLog/dao"
)

type Task struct {
	Id         uint `gorm:"primaryKey"`
	SessionId  string
	DeviceType int
	StartTime  int64
	EndTime    int64
	TaskId     string `gorm:"unique"`
	FileName   string
	State      int
	Msg        string
}

//
// GetSessionId
//  @Description: 根据TaskId获取SessionId
//  @param taskId
//  @return string
//  @return error
//
func GetSessionId(taskId string) (string, error) {
	task := Task{}
	if err := dao.DB.Table("task").Where("task_id = ?", taskId).Select("session_id").First(&task).Error; err != nil {
		return "", err
	}
	return task.SessionId, nil
}
