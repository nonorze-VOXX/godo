package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBcon() {
	var err error
	dbName := "./godo.db"
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Fatal("connect db fail.")
	}
}
