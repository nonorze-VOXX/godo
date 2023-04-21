package main

import (
	"godo/initializers"
	"godo/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.DBcon()
}
func main() {
	initializers.DB.AutoMigrate(&models.UnitTodo{})
}
