package models

import (
	"time"

	"github.com/JarGad23/go-rest-api/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, userId) 
		VALUES(?, ?, ?, ?, ?)
	`
	sqlStmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer sqlStmt.Close()

	result, err := sqlStmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.Id = id

	return err
}

func GetAllEvents() []Event {
	return events
}
