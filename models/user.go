package models

import (
	"errors"
	"ginmx2/db"
	"ginmx2/utils"
)

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
	hashpass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := data.Exec(u.Email, hashpass, u.Firstname, u.Lastname)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId() // get the last inserted id performed by data.Exec() and store it in the userId variable
	u.ID = userId                        // set the user id to the userId variable we got from the last inserted id
	return err
}

func (u User) ValidateCredentials() error {
	var queriedPassword string

	query := `SELECT id, password FROM users WHERE email = ?`
	data := db.DB.QueryRow(query, u.Email) // Passing query and specifying u.Email as the parameter to the query
	err := data.Scan(&queriedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}
	passwordCheck := utils.VerifyPassword(queriedPassword, u.Password)
	if !passwordCheck {
		return errors.New("invalid credentials")
	}
	return nil
}
