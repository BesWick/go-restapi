package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to goLang",
		Description: "how to create a Rest API",
	},
}

//create event
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter the event you'd like to add"+
			"along with its title and description")

	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)

}

//get one event
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["ID"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//get all events
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)

}

//update one event
func updateOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["ID"]

	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter")
	}

	json.Unmarshal(reqBody, &newEvent)

	for i, event := range events {
		if event.ID == eventID {
			event.Title = newEvent.Title
			event.Description = newEvent.Description
			events = append(events[:i], newEvent)
			json.NewEncoder(w).Encode(newEvent)
		}
	}

}

//delete one event
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["ID"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{ID}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{ID}", updateOneEvent).Methods("PATCH")
	router.HandleFunc("/events/{ID}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
