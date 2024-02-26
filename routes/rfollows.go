package routes

import (
	"ginmx2/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func followNote(context *gin.Context) {
	userId := context.GetInt64("userId")
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		return
	}
	note, err := models.GetNoteById(noteId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch note"})
		return
	}
	err = note.Follow(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow note"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Followed note"})
}

func unfollowNote(context *gin.Context) {
	userId := context.GetInt64("userId")
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var note models.Note
	note.ID = noteId
	err = note.Unfollow(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow note"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Unfollowed note"})
}
