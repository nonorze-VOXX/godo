package models

import (
	"time"

	"gorm.io/gorm"
)

type UnitTodo struct {
	gorm.Model
	Title string
	Stime time.Time
	Etime time.Time
	Done  bool
}
