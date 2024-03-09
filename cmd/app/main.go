package main

import (
	"ginmx2/db"
	"ginmx2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	gin.Default()
	server := gin.Default()
	server.LoadHTMLGlob("static/*")
	routes.BuildRoutes(server)
	err := server.Run(":3000")
	if err != nil {
		panic(err)
	}
}
