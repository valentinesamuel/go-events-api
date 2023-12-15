package models

import (
	"time"

	"example.com/events-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	StartDate   time.Time `binding:"required"`
	EndDate     time.Time `binding:"required"`
	UserId      int
}

var events []Event = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, startDate, endDate, userId)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.StartDate, e.EndDate, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() // Get the ID of the inserted row
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
