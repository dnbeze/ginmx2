package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database" + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	// Create users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    		id INTEGER PRIMARY KEY AUTOINCREMENT,
	    		email TEXT NOT NULL UNIQUE, 
	    		firstname TEXT NOT NULL,
	    		lastname TEXT NOT NULL,
	    		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Issue creating users table")
	}

	createNotesTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		customer TEXT NOT NULL,
		user_id INTEGER,
	    FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createNotesTable)
	if err != nil {
		panic("Issue creating notes table")
	}
}
