package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Resources
// https://medium.com/@hugo.bjarred/rest-api-with-golang-and-mux-e934f581b8b5
// https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work

type LunchSpot struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

var lunchSpots []LunchSpot

func main() {
	fmt.Println("Start")
	router := mux.NewRouter()

	// Create new lunchSpots for testing
	spot1 := LunchSpot{"1", "Shalom Y'all", "Delicious Israeli food", 3}
	spot2 := LunchSpot{"1", "Kati Thai", "Vegetarian Thai food", 5}
	lunchSpots = append(lunchSpots, spot1, spot2)

	router.HandleFunc("/suggestions", suggestions).Methods("POST")
	router.HandleFunc("/suggestions", getSuggestions).Methods("GET")

	fmt.Println("Listen and Server")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(corsObj)(router)))
}

func suggestions(rw http.ResponseWriter, rq *http.Request) {

}

func getSuggestions(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(lunchSpots)
}
