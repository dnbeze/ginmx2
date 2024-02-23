package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // generate hashed password
	return string(bytes), err                                                       // return hashed password as string and error
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) // compare hashed password with password
	return err == nil                                                              // boool if hashed password and password are the same
}
