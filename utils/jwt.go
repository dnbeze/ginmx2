package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secretkey"

func GenerateJWT(email string, userId int64) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"userId":    userId,
		"expiresAt": time.Now().Add(time.Minute * 30).Unix(),
	})
	return tok.SignedString(secretKey) // return the signed token
}
