package object

import (
	"time"

	"github.com/aecra/covid/db"
)

type Record struct {
	ID           int       `gorm:"primaryKey auto_increment"`
	Time         time.Time `gorm:"autoCreateTime" json:"time"`
	Username     string    `gorm:"not null" json:"username"`
	Email        string    `gorm:"not null" json:"email"`
	Position     string    `gorm:"default:'school'" json:"position"`
	ReportResult string    `gorm:"default:''" json:"report_result"`
	NoticeResult string    `gorm:"default:''" json:"notice_result"`
}

func AddRecord(record *Record) {
	db.GetConnection().Create(record)
}

func GetRecords(user *User) []Record {
	var records []Record
	db.GetConnection().Where("username = ?", user.Username).Order("time desc").Limit(50).Find(&records)
	return records
}
