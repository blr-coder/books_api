package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blr-coder/books_api/controllers"
	"github.com/blr-coder/books_api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong!"})
	})

	r.POST("/books", controllers.CreateBook)

	r.Run()
}
