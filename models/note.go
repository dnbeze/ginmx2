package models

import (
	"ginmx2/db"
	"time"
)

type Note struct {
	ID       int64
	Title    string    `binding:"required"`
	Body     string    `binding:"required"`
	DateTime time.Time `binding:"required"`
	Customer string    `binding:"required"`
	UserID   int
}

func (n Note) Save() error {
	query := `
	INSERT INTO notes (title, body, date_time, customer, user_id) 
	VALUES (?, ?, ?, ?, ?)` // ? is a placeholder for the actual value provides sql injection protection apparently
	data, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer data.Close()
	result, err := data.Exec(n.Title, n.Body, n.DateTime, n.Customer, n.UserID) // If you have a query to change things use Exec, if you have a query to get things use Query
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	n.ID = id
	return err
}

func GetAllNotes() ([]Note, error) {
	query := `SELECT * FROM notes`
	data, err := db.DB.Query(query) // if you a query to fetch data use Query
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var notes []Note  // create a slice of notes to store Note(s) returned from the database
	for data.Next() { // iterate through the data
		var note Note // create a new note of type Note
		err := data.Scan(&note.ID, &note.Title, &note.Body, &note.DateTime, &note.Customer, &note.UserID)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note) // append to the notes slice
	}
	return notes, err
}

func GetNoteById(id int64) (*Note, error) { // *Note is a pointer to a Note for error handling purposes
	query := `SELECT * FROM notes WHERE id = ?`
	data := db.DB.QueryRow(query, id) // if you a query to fetch data use Query
	var note Note
	err := data.Scan(&note.ID, &note.Title, &note.Body, &note.DateTime, &note.Customer, &note.UserID)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (n Note) UpdateNote() error {
	query := `
	UPDATE notes 
	SET title = ?, body = ?, date_time = ?, customer = ? 
	WHERE id = ? 
	` // WHERE id = ? is a condition to update the note with the id provided
	data, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer data.Close()
	_, err = data.Exec(n.Title, n.Body, n.DateTime, n.Customer, n.ID)
	return err
}
