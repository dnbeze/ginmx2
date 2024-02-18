package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.Default()
	server := gin.Default()

	server.GET("/ping", getPong)

	server.Run(":3000")
}

func getPong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
		"message": "pong",
	})

}
