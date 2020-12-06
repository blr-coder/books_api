package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blr-coder/books_api/models"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book = models.Book{Title: book.Title, Author: book.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
