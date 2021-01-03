package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/auth"
)

func Authenticate(ctx *gin.Context) {
	logrus.Info("Auth")
	jwtToken, err := auth.GenerateJWT()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token generate error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": jwtToken})
}
