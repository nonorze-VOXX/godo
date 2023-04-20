package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Dbcon() {
	var err error

	dbName := "./godo.db"
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	//
	// DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal("fail connect database")
	}

}
