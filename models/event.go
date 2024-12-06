package models

import (
	"fmt"

	"github.com/KarasunoAs9/RestAPI-Project/db"
)



type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int64
}

//var event *Event



func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, user_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.UserID).Scan(&e.ID)

	if err != nil {
		return fmt.Errorf("error with save data: %v", err)
	}

	return err
}

func (e *Event) GetEventById(id int64) error {
	query := `SELECT * FROM events WHERE id = $1`

	row := db.DB.QueryRow(query, id)

	if row == nil {
		return fmt.Errorf("no such event")
	}

	return nil
}