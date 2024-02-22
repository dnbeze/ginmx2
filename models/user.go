package models

import "ginmx2/db"

type User struct {
	ID        int64
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Firstname string `binding:"required"`
	Lastname  string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password, firstname, lastname) 
	VALUES (?, ?, ?, ?)`

	data, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer data.Close()
	result, err := data.Exec(u.Email, u.Password, u.Firstname, u.Lastname)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId() // get the last inserted id performed by data.Exec() and store it in the userId variable
	u.ID = userId                        // set the user id to the userId variable we got from the last inserted id
	return err
}
