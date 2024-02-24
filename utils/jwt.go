package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secretkey" //TODO fix secret key to something secure

func GenerateJWT(email string, userId int64) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"userId":    userId,
		"expiresAt": time.Now().Add(time.Minute * 30).Unix(),
	})
	return tok.SignedString([]byte(secretKey)) // return the signed token signed with the secret key
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // parse the token and return the claims
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // check if the token is signed with the HMAC method and return bool ok

		if !ok {
			return nil, jwt.ErrSignatureInvalid // return error if the token is not signed with the HMAC method
		}
		return secretKey, nil
	})
}
