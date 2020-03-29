// This is the main package
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", routes.homeLink)
	router.HandleFunc("/event", routes.createEvent).Methods("POST")
	router.HandleFunc("/event/{id}", routes.getOneEvent).Methods("GET")
	router.HandleFunc("/events", routes.getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", routes.updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", routes.deleteEvent).Methods("DELETE")

	eventDB, _ := event.OpenEventDB()
	defer eventDB.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
