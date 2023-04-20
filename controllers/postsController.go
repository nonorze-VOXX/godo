package controllers

import (
	"godo/initializers"
	"godo/models"

	"github.com/gin-gonic/gin"
)

func GetReturn(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func Postfunction(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": "pongpong",
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	initializers.DB.Find(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}
func Update(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	c.JSON(200, gin.H{
		"posts": post,
	})
}
