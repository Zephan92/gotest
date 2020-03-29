// Package event contains all event functionality
package event

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

// Event struct that contains data related to a
type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []Event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

var db *sql.DB

// OpenEventDB opens a connection to the db and initials the events table
func OpenEventDB() (eventDB *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./goTest.db")
	if err != nil {
		log.Fatal(err)
	}
	initEventsDatabase()
	return db, nil
}

// AddEvent appends the specified event into the database
func AddEvent(event Event) (eventID int64) {
	statement, _ := db.Prepare("INSERT INTO event (title, description) VALUES (?, ?)")
	result, _ := statement.Exec(event.Title, event.Description)
	eventID, _ = result.LastInsertId()
	return eventID
}

// GetEvent querys for a specified event by id
func GetEvent(eventID int64) (event Event) {
	queryStr := fmt.Sprintf("SELECT id, title, description FROM event WHERE id = %d", eventID)
	rows, _ := db.Query(queryStr)

	for rows.Next() {
		event = Event{
			ID: strconv.FormatInt(int64(eventID), 10),
		}
		rows.Scan(&event.ID, &event.Title, &event.Description)
		break
	}

	return event
}

func initEventsDatabase() {
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS event (id INTEGER PRIMARY KEY, title TEXT, description TEXT)")
	statement.Exec()
}
