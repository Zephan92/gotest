// Package routes bundles all routes in one package
package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	// events = append(events, newEvent)
	eventID := event.AddEvent(newEvent)
	newEvent.ID = strconv.FormatInt(int64(eventID), 10)
	log.Println(fmt.Sprintf("Created new event '%d'", eventID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventIDStr := mux.Vars(r)["id"]
	eventID, _ := strconv.ParseInt(eventIDStr, 10, 64)
	event := event.GetEvent(eventID)
	json.NewEncoder(w).Encode(event)
}

// func getAllEvents(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(events)
// }

// func updateEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]
// 	var updatedEvent Event

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}
// 	json.Unmarshal(reqBody, &updatedEvent)

// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			singleEvent.Title = updatedEvent.Title
// 			singleEvent.Description = updatedEvent.Description
// 			events = append(events[:i], singleEvent)
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func deleteEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			events = append(events[:i], events[i+1:]...)
// 			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
// 		}
// 	}
// }
