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

	r.GET("/", controllers.GetMainPage)
	r.GET("/godoList", controllers.GetGodoList)
	r.Run()

}
