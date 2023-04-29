package main

import (
	"godo/controllers"
	"godo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.DBcon()
	initializers.LoadEnvVar()
}

func main() {
	r := gin.Default()

	r.GET("/godoList", controllers.GetGodoList)
	r.POST("/godoList", controllers.SubmitGodo)
	r.POST("/godoList/update", controllers.UpdateUnitGodo)
	r.Run()
}
