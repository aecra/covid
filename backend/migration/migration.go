package migration

import (
	"github.com/aecra/covid/db"
	"github.com/aecra/covid/object"
)

func AutoMigrate() {
	db.GetConnection().AutoMigrate(&object.User{})
	db.GetConnection().AutoMigrate(&object.Record{})
}
