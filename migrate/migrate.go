package main

import (
	"godo/initializers"
	"godo/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.Dbcon()
}

func main() {
	initializers.DB.AutoMigrate((&models.Post{}))
}
