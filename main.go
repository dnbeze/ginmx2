package main

import (
	"ginmx2/db"
	"ginmx2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	gin.Default()
	server := gin.Default()

	server.GET("/ping", getPong)
	server.GET("/notes", getNotes)
	server.POST("/notes", createNote)
	err := server.Run(":3000")
	if err != nil {
		panic(err)
	}
}

func getPong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"}) // gin.H is a shortcut for map[string]interface{} and is used to perform JSON response
}

func getNotes(context *gin.Context) {
	notes, err := models.GetAllNotes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes."})
		return
	}
	context.JSON(http.StatusOK, notes) // gin.H is a shortcut for map[string]interface{} and is used to perform JSON response
}

func createNote(context *gin.Context) {
	var note models.Note
	err := context.ShouldBindJSON(&note)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode JSON"})
		return
	}
	note.ID = 1
	note.UserID = 1

	err = note.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "note": note})
}
