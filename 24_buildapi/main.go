package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Team struct {
	TeamId    string `json:"tid"`
	TeamName  string `json:"tname"`
	TeamPrice int    `json:"price"`
	Hero      *Hero  `json:"hero"`
}

type Hero struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// =--=- FAke DB ==--=-//

var team []Team

// =--=-= Middleware, helper -=-==-//
func (t *Team) IsEmpty() bool {
	return t.TeamId == "" && t.TeamName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// =-=- serve home route -=-=-///
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Team API</h1>"))
}

func getAllTeam(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses");
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}