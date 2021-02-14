package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/auth"
	"github.com/blr-coder/books_api/database"
	"github.com/blr-coder/books_api/handlers"
)

func main() {
	router := gin.New()

	database.ConnectDatabase()

	// Test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong!"})
	})

	// Users
	router.POST("/register", handlers.RegisterUser)

	// Auth
	router.POST("/auth", handlers.Authenticate)
	router.POST("/test_token_parse", handlers.Parse)

	// Books api
	booksAPI := router.Group("/api", auth.Middleware)
	{
		booksAPI.POST("/books", auth.Middleware, handlers.CreateBook)
		// router.POST("/books", handlers.CreateBook)
		booksAPI.GET("/books", handlers.AllBooks)
		booksAPI.GET("/books/:id", handlers.GetBook)

		booksAPI.DELETE("/books/:id", handlers.DeleteBook)
	}

	err := router.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
	if err != nil {
		logrus.Error(err)
	}
}
