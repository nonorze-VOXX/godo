package main

import (
	"godo/controllers"
	"godo/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.DBcon()
	initializers.LoadEnvVar()
}

func main() {
	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	server.Use(cors.New(corsConfig))

	server.GET("/godoList", controllers.GetGodoList)
	server.POST("/godoList", controllers.SubmitGodo)
	server.POST("/godoList/update", controllers.UpdateUnitGodo)

	server.Run()
}
