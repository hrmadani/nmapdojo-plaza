package models

import (
	"time"

	"github.com/hrmadani/nmapdojo-plaza/pkg/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type UserReport struct {
	gorm.Model
	ReportId   int    `json:"report_id"`   //ziro is null
	ReportType string `json:"report_type"` //Police - Traffic Jam - Car Crash
	UserId     int    `json:"user_id"`
	Action     string `json:"action"` //add , like , dislike
	CreatedAt  time.Time
}

//All Report
type Report struct {
	gorm.Model
	ReportType int       `json:"report_type"`
	ExpireTime time.Time `gorm:"index" json:"expire_time"`
}

//Database connection
func init() {
	//Connect to database
	config.Connect()
	db = config.GetDB()
}

func (userReportLog UserReport) SetNowTime() UserReport {
	userReportLog.CreatedAt = time.Now()
	return userReportLog
}

//All alive reports get from database
func GetAllAliveReport() ([]Report, error) {
	//Get all alive reports
	var Report []Report
	db.Where("expire_time > ?", time.Now()).Find(&Report)
	return Report, nil
}
