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


func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, user_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.UserID).Scan(&e.ID)

	if err != nil {
		return fmt.Errorf("error with save data: %v", err)
	}

	return err
}

func (e *Event) UpdateEvent(id, userId int64) error {
	query := `
	UPDATE events
	SET name = $1,
			description = $2,
			location = $3
			WHERE id = $4 and user_id = $5;
	`

	row := db.DB.QueryRow(query, e.Name, e.Description, e.Location, id, userId)

	if row == nil {
		return fmt.Errorf("no such event on database")
	}

	return nil

}

func (e *Event) DeleteEvent(id, userId int64) error {
	query := "DELETE FROM events WHERE id = $1 AND user_id = $2"

	_, err := db.DB.Exec(query, id, userId)

	if err != nil {
		return fmt.Errorf("no such event on database")
	}
	
	return nil
}

func GetAllEvent(userId int64) ([]*Event, error) {
	
	var events []*Event
	query := `SELECT * FROM events WHERE user_id = $1`

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return nil, fmt.Errorf("error with select events: %v", err)
	}

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID); err != nil {
			return nil, fmt.Errorf("error with scan data: %v", err)
		}
		events = append(events, &event)
	}

	defer rows.Close()

	return events, nil
}