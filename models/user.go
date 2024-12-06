package models

import (
	"errors"

	"github.com/KarasunoAs9/RestAPI-Project/db"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID int64
	Username string
	UserPassword string
}



func (u *User) Save() error {
	query := `INSERT INTO users (username, userpassword) VALUES $1, $2`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.UserPassword), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("error with hash password")
	}

	_, err = db.DB.Exec(query, u.Username, hashedPassword)

	if err != nil {
		return errors.New("error with insert user to database")
	}

	return err
}