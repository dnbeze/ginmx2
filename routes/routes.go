package routes

import (
	"ginmx2/middleware"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(server *gin.Engine) {
	server.GET("/ping", getPong) // GET POST PUT PATCH DELETE OPTIONS HEAD
	server.GET("/notes", getNotes)
	server.GET("/notes/:id", getNote)                          // dynamic route
	server.POST("/notes", middleware.Authenticate, createNote) // NEW entry // protected // can register multiple request handlers for middleware
	server.PUT("/notes/:id", updateNote)                       // UPDATE entry
	server.DELETE("/notes/:id", deleteNote)                    // DELETE entry
	server.POST("/signup", signup)                             // new user
	server.POST("/login", login)                               // login
}
