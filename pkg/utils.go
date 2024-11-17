package pkg

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("")

func SetSecret(secret string) {
	jwtSecret = []byte(secret)
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
