package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/auth"
	"github.com/blr-coder/books_api/database"
	"github.com/blr-coder/books_api/models"
)

type authInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(ctx *gin.Context) {
	logrus.Info("Auth")

	var input authInput

	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	var user models.User

	err := database.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	jwtToken, err := auth.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token generate error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
