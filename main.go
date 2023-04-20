package main

import (
	"godo/controllers"
	"godo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.Dbcon()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.GetReturn)
	r.POST("/post", controllers.Postfunction)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)
	r.POST("/postup/:id", controllers.Update)
	r.Run()
}

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Print("amogus")
// 	}
// }
