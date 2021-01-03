package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const (
	signingKey = "super_secret_signing_key"
	tokenTTL   = 60 * time.Second
)

func GenerateJWT() (string, error) {
	// генерим новый токен со стандартными полями используя библиотеку jwt
	token := jwt.New(jwt.SigningMethodHS256)
	// получаем все поля токена в виде map
	claims := token.Claims.(jwt.MapClaims)
	// добавляем поля с данными пользователя
	claims["userId"] = 1        // для примера (в реальной ситуации мы должны передавать в метод данные юзера)
	claims["userName"] = "John" // для примера (в реальной ситуации мы должны передавать в метод данные юзера)
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
