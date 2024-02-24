package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secretkey" //TODO fix secret key to something secure

func GenerateJWT(email string, userId int64) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // create a new token with the HMAC signing method and the claims using the MapClaims type
		"email":     email,
		"userId":    userId,
		"expiresAt": time.Now().Add(time.Minute * 30).Unix(), // set the expiration time to 30 minutes from now
	})
	return tok.SignedString([]byte(secretKey)) // return the signed token signed with the secret key HMAC expects a byte slice, I convert the secret key to a byte slice
}

func VerifyJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(tokenString *jwt.Token) (interface{}, error) { // parse the token and return the claims. This is an anonymous function that takes a tokenString and returns an interface and an error
		_, ok := tokenString.Method.(*jwt.SigningMethodHMAC) // check if the token is signed with the HMAC method and return bool ok
		if !ok {
			return nil, jwt.ErrSignatureInvalid // return error if the token is not signed with the HMAC method
		}
		return []byte(secretKey), nil // return the secret key and nil error if the token is signed with the HMAC method to the Parse function
	}) //end of anonymous function and Parse function call
	if err != nil {
		return 0, jwt.ErrTokenUnverifiable
	}
	tokenIsValid := token.Valid // check if the token is valid
	if !tokenIsValid {          // if the token is not valid, return an error
		return 0, jwt.ErrTokenSignatureInvalid
	}
	//TODO move claim data extraction to its own function
	claims, ok := token.Claims.(jwt.MapClaims) // get the claims from the token and define it as a jwt.MapClaims because we use MapClaims type in GenerateJWT()
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}
	//email := claims["email"].(string)  // get the email from the claims and define it as a string
	userId := claims["userId"].(int64) // get the userId from the claims and define it as an int64
	return userId, nil
}
