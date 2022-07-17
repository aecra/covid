package object

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Init() {
	var db_dsn string
	if os.Getenv("DB_DSN") != "" {
		db_dsn = os.Getenv("DB_DSN")
	} else {
		db_dsn = "root:123456@tcp(127.0.0.1:3306)/covid"
	}
	db, err := gorm.Open(mysql.Open(db_dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Record{})
}
