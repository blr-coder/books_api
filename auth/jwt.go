package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/blr-coder/books_api/models"
)

const (
	signingKey = "super_secret_signing_key"
	tokenTTL   = 60 * time.Second
)

func GenerateJWT(user models.User) (string, error) {
	logrus.Info("User - ", user)
	// генерим новый токен со стандартными полями используя библиотеку jwt
	token := jwt.New(jwt.SigningMethodHS256)
	// получаем все поля токена в виде map
	claims := token.Claims.(jwt.MapClaims)
	// добавляем поля с данными пользователя
	claims["userId"] = user.ID
	claims["userEmail"] = user.Email
	// добавляем поле с временем жизни токена
	claims["tokenExpire"] = time.Now().Add(tokenTTL).Unix()
	// приводим к строке
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return tokenString, nil
}
