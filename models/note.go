package models

import "time"

type Note struct {
	ID       int       `json:"id"`
	Title    string    `binding:"required" json:"title"`
	Body     string    `binding:"required" json:"body"`
	DateTime time.Time `binding:"required"`
	UserID   int
}

var notes = []Note{}

func (n Note) Save() {
	// TODO save note to database
	notes = append(notes, n)
}

func GetAllNotes() []Note {
	return notes
}
