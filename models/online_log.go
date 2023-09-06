package models

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
