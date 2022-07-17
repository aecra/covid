package object

import "time"

type Record struct {
	ID       int       `gorm:"primaryKey auto_increment"`
	Time     time.Time `gorm:"autoCreateTime"`
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"not null"`
	Position string    `gorm:"default:'school'"`
	Content  string    `gorm:"default:''"`
	Result   string    `gorm:"default:''"`
}

func AddRecord(record *Record) {
	database.Create(record)
}
