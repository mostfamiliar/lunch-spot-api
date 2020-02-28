package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	Address     string `json:"address"`
	Link        string `json:"link"`
	Votes       int    `json:"votes"`
}

var lunchSpots []LunchSpot

func main() {
	fmt.Println("Start")
	router := mux.NewRouter()

	// Create new lunchSpots for testing
	spot1 := LunchSpot{"1", "Shalom Y'all", "Delicious Israeli food", "117 SE Taylor #101, Portland, OR 97214", "https://www.shalomyallpdx.com/", 3}
	spot2 := LunchSpot{"2", "Kati Thai", "Vegetarian Thai food", "2932 SE Division St, Portland, OR 97202", "http://www.katiportland.com/", 5}
	spot3 := LunchSpot{"3", "XLB", "Chinese comfort food and baozi", "4090 N Williams Ave, Portland, OR 97227", "https://www.xlbpdx.com/", 8}
	spot4 := LunchSpot{"4", "Por Que No?", "Tacos and bowls", "3524 N Mississippi Ave, Portland, OR 97227", "https://porquenotacos.com/", 3}
	lunchSpots = append(lunchSpots, spot1, spot2, spot3, spot4)

	router.HandleFunc("/suggestions", createSuggestions).Methods("POST")
	router.HandleFunc("/suggestions", getSuggestions).Methods("GET")

	fmt.Println("Listen and Server")
	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func createSuggestions(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var newLunchSpot LunchSpot
	_ = json.NewDecoder(rq.Body).Decode(&newLunchSpot)
	id := strconv.Itoa(len(lunchSpots) + 1)
	newLunchSpot.ID = id
	lunchSpots = append(lunchSpots, newLunchSpot)
	json.NewEncoder(rw).Encode(&newLunchSpot)
}

func getSuggestions(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(lunchSpots)
}
