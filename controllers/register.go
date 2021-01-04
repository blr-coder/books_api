package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/database"
	"github.com/blr-coder/books_api/models"
)

func RegisterUser(ctx *gin.Context) {
	logrus.Info("RegisterUser")
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user = models.User{Email: user.Email, Password: user.Password, Role: "user"}
	database.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
