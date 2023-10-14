package models

type DefaultOutputStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"err_msg"`
}

// LoginInputStruct 登录传入结构体
type LoginInputStruct struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// LoginOutputStruct 登录返回结构体
type LoginOutputStruct struct {
	DefaultOutputStruct
	Token *string `json:"token"`
}

// CreateAppInputStruct 创建应用传入结构体
type CreateAppInputStruct struct {
	AppName string `json:"app_name"`
}

// CreateAppOutputStruct 创建应用返回结构体
type CreateAppOutputStruct struct {
	DefaultOutputStruct
	AppId string `json:"app_id"`
}

// ExchangePubKeyInputStruct 交换公钥传入结构体
type ExchangePubKeyInputStruct struct {
	ClientPubKey  string `json:"client_pub_key"`
	ExpireSeconds int    `json:"expire_seconds"`
}

// ExchangePubKeyOutputStruct 交换公钥返回结构体
type ExchangePubKeyOutputStruct struct {
	DefaultOutputStruct
	TmpSessionId string `json:"tmp_session_id"`
	ServerPubKey string `json:"server_pub_key"`
}

// VerifySharedKeyInputStruct 验证共享密钥传入结构体
type VerifySharedKeyInputStruct struct {
	TmpSessionId string `json:"tmp_session_id"`
	VerifyData   string `json:"verify_data"`
}

// VerifySharedKeyOutputStruct 验证共享密钥返回结构体
type VerifySharedKeyOutputStruct struct {
	DefaultOutputStruct
	DecryptData string `json:"decrypt_data"`
}

// DeviceRegisterInputStruct 设备注册传入结构体
type DeviceRegisterInputStruct struct {
	AppId      string `json:"app_id"`
	DeviceType int    `json:"device_type"`
	DeviceName string `json:"device_name"`
	DeviceId   string `json:"device_id"`
}

// DeviceRegisterOutputStruct 设备注册返回结构体
type DeviceRegisterOutputStruct struct {
	DefaultOutputStruct
	SessionId string `json:"session_id"`
}

// PutOnlineLogInputStruct 上传实时日志传入结构体
type PutOnlineLogInputStruct struct {
	Data []PutOnlineLogInfoStruct `json:"data"`
}

// PutOnlineLogInfoStruct  实时日志结构体
type PutOnlineLogInfoStruct struct {
	Sequence      int64  `json:"sequence"`
	SystemVersion string `json:"system_version"`
	AppVersion    string `json:"app_version"`
	TimeStamp     int64  `json:"time_stamp"`
	LogLevel      int    `json:"log_level"`
	Identify      string `json:"identify"`
	Tag           string `json:"tag"`
	Msg           string `json:"msg"`
}

// GetTaskInputStruct 查询任务的传入结构体
type GetTaskInputStruct struct {
	DeviceType int `form:"device_type"`
}

// GetTaskOutputStruct 查询任务的返回结构体
type GetTaskOutputStruct struct {
	DefaultOutputStruct
	Data []GetTaskInfoStruct `json:"data"`
}

// GetTaskInfoStruct 任务结构体
type GetTaskInfoStruct struct {
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	TaskId    string `json:"task_id"`
}

// UploadLogFileErrCallBackInputStruct 日志无法上传时的反馈传入结构体
type UploadLogFileErrCallBackInputStruct struct {
	TaskId string `json:"task_id"`
	Msg    string `json:"msg"`
}

// GetAppListInputStruct 查询应用列表传入结构体
type GetAppListInputStruct struct {
	Page int `form:"page"`
}

// GetAppListOutputStruct 查询应用列表返回结构体
type GetAppListOutputStruct struct {
	DefaultOutputStruct
	Count int64                  `json:"count"`
	Data  []GetAppListInfoStruct `json:"data"`
}

// GetAppListInfoStruct 应用列表结构体
type GetAppListInfoStruct struct {
	AppId      string `json:"app_id"`
	AppName    string `json:"app_name"`
	CreateTime int64  `json:"create_time"`
}

// DeleteAppInputStruct 删除应用传入结构体
type DeleteAppInputStruct struct {
	AppId string `json:"app_id"`
}

// GetDeviceListInputStruct 查询应用下的设备列表传入结构体
type GetDeviceListInputStruct struct {
	AppId    string `form:"app_id"`
	Identify string `form:"identify"`
	Page     int    `form:"page"`
}

// GetDeviceListOutputStruct 查询应用下的设备列表返回结构体
type GetDeviceListOutputStruct struct {
	DefaultOutputStruct
	Count int64                     `json:"count"`
	Data  []GetDeviceListInfoStruct `json:"data"`
}

// GetDeviceListInfoStruct 设备列表结构体
type GetDeviceListInfoStruct struct {
	DeviceType int    `json:"device_type"`
	DeviceName string `json:"device_name"`
	DeviceId   string `json:"device_id"`
	SessionId  string `json:"session_id"`
}

// GetDeviceLogInputStruct 查询设备的日志传入结构体
type GetDeviceLogInputStruct struct {
	Page          int    `form:"page"`
	SessionId     string `form:"session_id"`
	SystemVersion string `form:"system_version"`
	AppVersion    string `form:"app_version"`
	StartStamp    int64  `form:"start_stamp"`
	EndStamp      int64  `form:"end_stamp"`
	LogLevel      int    `form:"log_level"`
	Identify      string `form:"identify"`
	Tag           string `form:"tag"`
	Msg           string `form:"msg"`
}

// GetDeviceLogOutputStruct 查询设备的日志返回结构体
type GetDeviceLogOutputStruct struct {
	DefaultOutputStruct
	Count int64                    `json:"count"`
	Data  []GetDeviceLogInfoStruct `json:"data"`
}

// GetDeviceLogInfoStruct 设备日志结构体
type GetDeviceLogInfoStruct struct {
	Sequence      int64  `json:"sequence"`
	SystemVersion string `json:"system_version"`
	AppVersion    string `json:"app_version"`
	TimeStamp     int64  `json:"time_stamp"`
	LogLevel      int    `json:"log_level"`
	Identify      string `json:"identify"`
	Tag           string `json:"tag"`
	Msg           string `json:"msg"`
}

// GetAllTaskInputStruct 查询任务列表传入结构体
type GetAllTaskInputStruct struct {
	AppId   string `form:"app_id"`
	TaskDes string `form:"task_des"`
	Page    int    `form:"page"`
}

// GetAllTaskOutputStruct 查询任务列表返回结构体
type GetAllTaskOutputStruct struct {
	DefaultOutputStruct
	Count int64                  `json:"count"`
	Data  []GetAllTaskInfoStruct `json:"data"`
}

// GetAllTaskInfoStruct 任务列表结构体
type GetAllTaskInfoStruct struct {
	TaskDes    string `json:"task_des"`
	SessionId  string `json:"session_id"`
	DeviceType int    `json:"device_type"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	TaskId     string `json:"task_id"`
	State      int    `json:"state"`
	Msg        string `json:"msg"`
}

// CreateTaskInputStruct 创建任务传入结构体
type CreateTaskInputStruct struct {
	AppId      string `json:"app_id"`
	TaskDes    string `json:"task_des"`
	SessionId  string `json:"session_id"`
	DeviceType int    `json:"device_type"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
}

// CreateTaskOutputStruct 创建任务返回结构体
type CreateTaskOutputStruct struct {
	DefaultOutputStruct
	TaskId string `json:"task_id"`
}

// DeleteTaskInputStruct 删除任务传入结构体
type DeleteTaskInputStruct struct {
	TaskId string `json:"task_id"`
}

// GetTaskLogInputStruct 查询任务日志传入结构体
type GetTaskLogInputStruct struct {
	Page          int    `form:"page"`
	TaskId        string `form:"task_id"`
	SystemVersion string `form:"system_version"`
	AppVersion    string `form:"app_version"`
	StartStamp    int64  `form:"start_stamp"`
	EndStamp      int64  `form:"end_stamp"`
	LogLevel      int    `form:"log_level"`
	Identify      string `form:"identify"`
	Tag           string `form:"tag"`
	Msg           string `form:"msg"`
}

// GetTaskLogOutputStruct 查询任务日志返回结构体
type GetTaskLogOutputStruct struct {
	DefaultOutputStruct
	Count int64                  `json:"count"`
	Data  []GetTaskLogInfoStruct `json:"data"`
}

// GetTaskLogInfoStruct 任务日志结构体
type GetTaskLogInfoStruct struct {
	Sequence      int64  `json:"sequence"`
	SystemVersion string `json:"system_version"`
	AppVersion    string `json:"app_version"`
	TimeStamp     int64  `json:"time_stamp"`
	LogLevel      int    `json:"log_level"`
	Identify      string `json:"identify"`
	Tag           string `json:"tag"`
	Msg           string `json:"msg"`
}
