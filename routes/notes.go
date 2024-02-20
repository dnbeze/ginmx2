package routes

import (
	"ginmx2/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func getNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id from the dynamic route /notes/:id and store it in the id variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		return
	}
	note, err := models.GetNoteById(noteId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch note."})
		return
	}
	context.JSON(http.StatusOK, note)

}
