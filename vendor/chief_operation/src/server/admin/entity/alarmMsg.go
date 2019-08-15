package entity

import "time"



type AlarmMsg struct {
	Id          int64    `gorm:"primary_key;column:id"`
	AppName        string    `gorm:"column:appName"`
	Org  string     `gorm:"column:org"`
	StartTime     time.Time  `gorm:"column:startTime"`
	CreateTime     time.Time  `gorm:"column:createTime"`
	Title string   `gorm:"column:title"`
	ErrMsg string   `gorm:"column:errMsg"`
	ModuleName string   `gorm:"column:moduleName"`
}

func (this *AlarmMsg) TableName() string {
	return "alarm_msg"
}