package routes

import "github.com/gin-gonic/gin"

func BuildRoutes(server *gin.Engine) {
	server.GET("/ping", getPong)
	server.GET("/notes", getNotes)
	server.GET("/notes/:id", getNote) // dynamic route
	server.POST("/notes", createNote)

}
