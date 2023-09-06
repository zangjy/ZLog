package models

type Task struct {
	Id         uint `gorm:"primaryKey"`
	SessionId  string
	DeviceType int
	StartTime  int64
	EndTime    int64
	TaskId     string `gorm:"unique"`
	FileName   string
	State      int
}
