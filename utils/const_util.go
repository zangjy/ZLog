package utils

// EncryptingKey Token加密密钥
const EncryptingKey = "5dfjNGwIO4Kt5C2WcS1qsApGb3c8DCyd"

const (
	StatusNotFoundCode = "404"
	SuccessCode        = "0000"
	ErrorCode          = "0001"
	SessionId          = "SESSION_ID"
	TmpSessionId       = "TMP_SESSION_ID"
	Token              = "TOKEN"
	LogFileRootPath    = "./static/zlog"
)

const (
	V1Path                   = "/api/v1"
	LoginPath                = "/login"
	ExchangePubKeyPath       = "/exchange_pub_key"
	VerifySharedKeyPath      = "/verify_shared_key"
	DeviceRegisterPath       = "/device_register"
	CreateAppPath            = "/create_app"
	DeleteAppPath            = "/delete_app"
	PutOnlineLogPath         = "/put_online_log"
	GetTaskPath              = "/get_task"
	UploadLogFilePath        = "/upload_log_file"
	UploadLogFileErrCallBack = "/upload_log_file_err_callback"
	GetAppListPath           = "/get_app_list"
	GetDeviceListPath        = "/get_device_list"
	GetDeviceLogPath         = "/get_device_log"
	GetAllTaskPath           = "/get_all_task"
	CreateTaskPath           = "/create_task"
	DeleteTaskPath           = "/delete_task"
	GetTaskLogPath           = "/get_task_log"
)
