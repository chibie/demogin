package main

import (
	"demogin/controllers"
	"demogin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
