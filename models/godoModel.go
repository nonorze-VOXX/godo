package models

import (
	"time"

	"gorm.io/gorm"
)

type UnitTodo struct {
	gorm.Model
	title string
	sTime time.Time
	eTime time.Time
}
