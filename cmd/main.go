package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/controllers"
	"github.com/blr-coder/books_api/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	// Test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong!"})
	})

	// Books
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.AllBooks)
	router.GET("/books/:id", controllers.GetBook)

	router.DELETE("/books/:id", controllers.DeleteBook)

	err := router.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
	if err != nil {
		logrus.Error(err)
	}
}
