package models

import "time"

type Event struct {
	ID          int      
	Name       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	StartDate   time.Time `binding:"required"`
	EndDate     time.Time `binding:"required"`
	UserId      int       
}

var events []Event = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
