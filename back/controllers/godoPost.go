package controllers

import (
	"fmt"
	"godo/initializers"
	"godo/models"
	"time"

	"github.com/gin-gonic/gin"
)

func SubmitGodo(c *gin.Context) {
	var body struct {
		Title string
		Stime time.Time
		Etime time.Time
		Done  bool
	}

	c.Bind(&body)

	unitGodo := models.UnitTodo{
		Title: body.Title,
		Stime: body.Stime,
		Etime: body.Etime,
		Done:  body.Done,
	}
	fmt.Println(unitGodo)
	result := initializers.DB.Create(&unitGodo)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{"resuklt": "Success"})
	// {
	//   "Title": "amogus",
	//   "Stime": "2018-12-10T13:45:00.000Z",
	//   "Etime": "2018-12-10T13:45:00.000Z"
	// }
}
func GetGodoList(c *gin.Context) {
	var posts []models.UnitTodo
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"godoList": posts,
	})
}
func UpdateUnitGodo(c *gin.Context) {
	var body struct {
		Id    int
		Title string
		Stime time.Time
		Etime time.Time
		Done  bool
	}

	c.Bind(&body)
	var post models.UnitTodo
	initializers.DB.First(&post, body.Id)
	initializers.DB.Model(&post).Updates(models.UnitTodo{
		Title: body.Title,
		Stime: body.Stime,
		Etime: body.Etime,
		Done:  body.Done,
	})
	c.JSON(200, gin.H{
		"update": "succc",
	})
}
