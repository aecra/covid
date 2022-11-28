package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func init() {
	var db_type = os.Getenv("COVID_DB_TYPE")
	if db_type == "" {
		db_type = "sqlite"
	}
	var db_dsn = os.Getenv("COVID_DB_DSN")
	if db_dsn == "" {
		db_dsn = "covid.db"
	}
	var err error
	// support mysql/postgres/sqlite/sqlserver
	switch db_type {
	case "mysql":
		database, err = gorm.Open(mysql.Open(db_dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("mysql connected")
	case "postgres":
		database, err = gorm.Open(postgres.Open(db_dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("postgres connected")
	case "sqlite":
		database, err = gorm.Open(sqlite.Open(db_dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("sqlite connected")
	case "sqlserver":
		database, err = gorm.Open(sqlserver.Open(db_dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("sqlserver connected")
	default:
		log.Fatal("unknown db type: " + db_type)
	}
}

func GetConnection() *gorm.DB {
	return database
}
