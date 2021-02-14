package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/models"
)

const (
	SigningKey = "super_secret_signing_key"
	tokenTTL   = 60 * time.Second
)

// Генерит новый токен
func GenerateJWT(user models.User) (string, error) {
	// генерим новый токен со стандартными полями используя библиотеку jwt
	token := jwt.New(jwt.SigningMethodHS256)
	// получаем все поля токена в виде map
	claims := token.Claims.(jwt.MapClaims)
	// добавляем в полученную мапу поля с данными пользователя
	claims["userId"] = user.ID
	claims["userEmail"] = user.Email
	claims["userRole"] = user.Role
	// добавляем поле с временем жизни токена
	claims["tokenExpire"] = time.Now().Add(tokenTTL).Unix()
	logrus.Info("Claims - ", token.Claims)
	// приводим к строке
	tokenString, err := token.SignedString([]byte(SigningKey))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return tokenString, nil
}

// Парсит токен и возвращает claims
func ParseJWT(accessToken string, signingKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		logrus.Info("token.c - ", token.Claims)
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Middleware
func Middleware(ctx *gin.Context) {
	logrus.Info("Middleware")
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if headerParts[0] != "bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := ParseJWT(headerParts[1], []byte(SigningKey))
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// добавляем в контекст claims нашего пользователя
	ctx.Set("userClaims", claims)
	logrus.Info("ctx - ", ctx)
}
