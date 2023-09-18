package models

import (
	"ZLog/dao"
	"sync"
)

type Task struct {
	Id         uint `gorm:"primaryKey"`
	SessionId  string
	DeviceType int
	StartTime  int64
	EndTime    int64
	TaskId     string `gorm:"unique"`
	State      int    `gorm:"default:0;comment:-1:客户端无法发送日志文件，可能是没有日志文件等原因，此时需要客户端需要将原因提交到msg字段里 0:客户端未响应 1:已成功接受到日志文件"`
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

//
// GetTaskList
//  @Description: 查询任务列表
//  @param sessionId
//  @param deviceType
//  @return []GetTaskInfoStruct
//
func GetTaskList(sessionId string, deviceType int) []GetTaskInfoStruct {
	var result []GetTaskInfoStruct
	dao.DB.Table("task").Where("session_id = ? AND device_type = ? AND state = 0", sessionId, deviceType).Find(&result)
	if result == nil {
		result = make([]GetTaskInfoStruct, 0)
	}
	return result
}

//
// NotifyTaskState
//  @Description: 更改任务状态
//  @param sessionId
//  @param taskId
//  @param state
//  @return bool
//  @return string
//
func NotifyTaskState(sessionId, taskId string, state int) (bool, string) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	if curState, err := getTaskState(sessionId, taskId); err != nil {
		return false, "查询任务状态失败，请检查SessionId和TaskId是否对应"
	} else if curState == 1 {
		return false, "该任务已处于完成状态"
	} else if err := dao.DB.Table("task").Where("session_id = ? AND task_id = ?", sessionId, taskId).Update("state", state).Error; err != nil {
		return false, "修改任务状态失败"
	}
	return true, ""
}

//
// NotifyTaskMsg
//  @Description: 修改
//  @param sessionId
//  @param taskId
//  @param msg
//  @return bool
//  @return error
//
func NotifyTaskMsg(sessionId, taskId, msg string) (bool, string) {
	updateData := map[string]interface{}{
		"state": -1,
		"msg":   msg,
	}
	if curState, err := getTaskState(sessionId, taskId); err != nil {
		return false, "查询任务状态失败，请检查SessionId和TaskId是否对应"
	} else if curState == 1 {
		return false, "该任务已处于完成状态"
	} else if err := dao.DB.Table("task").Where("session_id = ? AND task_id = ?", sessionId, taskId).Updates(updateData).Error; err != nil {
		return false, "修改失败"
	}
	return true, ""
}

//
// getTaskState
//  @Description: 查询任务状态
//  @param sessionId
//  @param taskId
//  @return int
//  @return error
//
func getTaskState(sessionId, taskId string) (int, error) {
	var state int
	err := dao.DB.Table("task").Where("session_id = ? And task_id = ?", sessionId, taskId).Select("state").Scan(&state).Error
	if err != nil {
		return 0, err
	}
	return state, nil
}
