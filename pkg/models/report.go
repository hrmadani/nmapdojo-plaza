package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//The struct of report types includes the Police, Traffic Jam, and Car crash reports
type ReportTypes struct {
	gorm.Model
	ID             int
	Type           string `json:"report_type"`                  //Police - Traffic Jam - Car Crash
	TypeID         int    `gorm:"unique" json:"report_type_id"` //Police == 1 - Traffic Jam == 2 - Car Crash == 3
	LifeSpan       int    `json:"life_span"`
	FeedbackEffect int    `json:"feedback_effect"`
}

//Every reports is logged
type ReportLog struct {
	gorm.Model
	ID           int       `json:"id"`
	ReportTypeId int       `json:"report_type_id"`
	UserId       int       `json:"user_id"`
	ExpireTime   time.Time `gorm:"index" json:"expire_time"`
	CreatedAt    time.Time
}

//Database connection
func init() {
	//Todo: Connect to database
}

//All alive reports get from database
func GetAllAliveReportLogs() ([]ReportLog, error) {
	//Todo: Get all alive reports
	return nil, nil
}
