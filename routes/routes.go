package routes

import "github.com/gin-gonic/gin"

func BuildRoutes(server *gin.Engine) {
	server.GET("/ping", getPong) // GET POST PUT PATCH DELETE OPTIONS HEAD
	server.GET("/notes", getNotes)
	server.GET("/notes/:id", getNote)       // dynamic route
	server.POST("/notes", createNote)       // NEW entry
	server.PUT("/notes/:id", updateNote)    // UPDATE entry
	server.DELETE("/notes/:id", deleteNote) // DELETE entry
	server.POST("/signup", signup)          // new user
}
