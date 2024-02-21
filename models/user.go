package models

import "ginmx2/db"

type User struct {
	ID        int64
	Username  string `binding:"required"`
	Password  string `binding:"required"`
	email     string `binding:"required"`
	firstname string `binding:"required"`
	lastname  string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (username, password, email, firstname, lastname) 
	VALUES (?, ?, ?, ?, ?)`

	data, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer data.Close()
	result, err := data.Exec(u.Username, u.Password, u.email, u.firstname, u.lastname)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId() // get the last inserted id performed by data.Exec() and store it in the userId variable
	u.ID = userId                        // set the user id to the userId variable we got from the last inserted id
	return err
}
