package middleware

import (
	"errors"
	"ginmx2/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("authorization") // get the token from the request header
	if token == "" {                                     // if the token is empty
		context.AbortWithError(http.StatusUnauthorized, errors.New("Token empty")) // abort the request and send an error response
		return
	}

	userId, err := utils.VerifyJWT(token) // verify the token

	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, errors.New("failed VerifyJWT"))
		return
	}
	context.Set("userId", userId) // set the user id in the context
	context.Next()
}
