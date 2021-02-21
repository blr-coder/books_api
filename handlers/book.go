package handlers

import (
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/database"
	"github.com/blr-coder/books_api/models"
)

type createBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`

	/*Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`*/
}

type updateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (i createBookInput) Validate() error {
	if i.Author == "" || i.Title == "" {
		return errors.New("all fields are required")
	}

	if utf8.RuneCountInString(i.Title) > 256 || utf8.RuneCountInString(i.Author) > 256 {
		return errors.New("the number of letters cannot be more than 256")
	}

	return nil
}

func CreateBook(ctx *gin.Context) {
	logrus.Info("CreateBook")

	var input createBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	database.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"new_book": book})
}

func AllBooks(ctx *gin.Context) {
	logrus.Info("AllBooks")
	var books []models.Book
	database.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(ctx *gin.Context) {
	logrus.Info("GetBook")
	var book models.Book

	err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context) {
	logrus.Info("UpdateBook")
	var book models.Book
	err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	// Validate input
	var input updateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("Validate OK")

	book.Title = input.Title
	book.Author = input.Author

	database.DB.Updates(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})

}

func DeleteBook(ctx *gin.Context) {
	logrus.Info("DeleteBook")
	var book models.Book

	err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	database.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
