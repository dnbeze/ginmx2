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
	err := context.ShouldBindJSON(&note) // bind json from the request body to note struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode JSON"})
		return
	}
	userId := context.GetInt64("userId") // get the user id from the context
	note.UserID = userId                 // set the user id of the note to the user id of the user who created the note
	err = note.Save()                    // attempt to save the note to the database
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

func updateNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id from the dynamic route /notes/:id and store it in the id variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		return
	}
	userId := context.GetInt64("userId")    // get the user id from the context
	note, err := models.GetNoteById(noteId) // get note by id
	if note.UserID != userId {              // if the user id of the note does not match the user id of the user who created the note
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch note."})
		return
	}
	var updatedNote models.Note
	err = context.ShouldBindJSON(&updatedNote)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode JSON"})
		return
	}
	updatedNote.ID = noteId // set the id of the updated note to the id of the note being updated set at the beginning
	err = updatedNote.UpdateNote()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Note updated successfully", "note": updatedNote})
}

func deleteNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id from the dynamic route /notes/:id and store it in the id variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		return
	}
	userId := context.GetInt64("userId")    // get the user id from the context
	note, err := models.GetNoteById(noteId) // get note by id
	if note.UserID != userId {              // if the user id of the note does not match the user id of the user who created the note
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = note.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
