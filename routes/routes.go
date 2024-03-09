package routes

import (
	"ginmx2/middleware"
	"ginmx2/templates"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(server *gin.Engine) {
	server.GET("/ping", getPong) // GET POST PUT PATCH DELETE OPTIONS HEAD
	server.GET("/notes", getNotes)
	server.GET("/notes/:id", getNote) // dynamic route
	server.GET("/", index)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)              // middleware defined for group
	authenticated.POST("/notes", createNote)                // NEW entry // protected // can register multiple request handlers for middleware
	authenticated.PUT("/notes/:id", updateNote)             // UPDATE entry
	authenticated.DELETE("/notes/:id", deleteNote)          // DELETE entry
	authenticated.POST("/notes/:id/follow", followNote)     // register for a note (new entry in another table)
	authenticated.DELETE("/notes/:id/follow", unfollowNote) // unregister for a note (delete entry in another table)

	server.POST("/signup", signup) // new user
	server.POST("/login", login)   // login

}

func index(context *gin.Context) {
	err := templates.Hello("BOB").Render(context, context.Writer)
	if err != nil {
		panic(err)
	}
}
