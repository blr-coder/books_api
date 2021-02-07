package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/models"
)

const (
	SigningKey = "super_secret_signing_key"
	tokenTTL   = 60 * time.Second
)

type Claims struct {
	jwt.StandardClaims
	UserEmail string `json:"user_email"`
}

// Генерит новый токен
func GenerateJWT(user models.User) (string, error) {
	logrus.Info("User - ", user)
	// генерим новый токен со стандартными полями используя библиотеку jwt
	token := jwt.New(jwt.SigningMethodHS256)
	// получаем все поля токена в виде map
	claims := token.Claims.(jwt.MapClaims)
	// добавляем поля с данными пользователя
	claims["userId"] = user.ID
	claims["UserEmail"] = user.Email
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

// Парсит токен
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
