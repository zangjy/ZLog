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
