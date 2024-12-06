package models

import (
	"errors"
	"fmt"

	"github.com/KarasunoAs9/RestAPI-Project/db"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID int64
	Username string
	UserPassword string
}



func (u *User) Save() error {
	query := `INSERT INTO users (username, userpassword) VALUES ($1, $2)`

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


func (u *User) VerifyUser() error {
	query := `SELECT id, userpassword FROM users WHERE username = $1 `

	row := db.DB.QueryRow(query, u.Username)

	var retrivedPassword string

	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return fmt.Errorf("error with find user in db: %v", err)
	}

	

	err = bcrypt.CompareHashAndPassword([]byte(retrivedPassword), []byte(u.UserPassword))

	if err != nil {
		return fmt.Errorf("uncorrect password")
	}
	
	return nil
}