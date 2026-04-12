package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
var DBteam []Team

// =--=-= Middleware, helper -=-==-//
func (t *Team) IsEmpty() bool {
	// return t.TeamId == "" && t.TeamName == ""
	return t.TeamName == ""
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
	json.NewEncoder(w).Encode(DBteam)
}

func getEachTeam(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get Each Team")
	w.Header().Set("Content-Type", "application/json")

	params :=mux.Vars(r)

	for _, t := range DBteam{
		if t.TeamId == params["id"]{
			json.NewEncoder(w).Encode(t)
			return;
		}
	}
	json.NewEncoder(w).Encode("No Team found with given id")
	return;
}

func createTeam(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create new Team")
	w.Header().Set("Content-Type", "application/json")

	// =-=-=-- if body is empty -=-=-=--//
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
		return;
	}	


	// =-=-=-- what if data is in wrong format like {} -=-=-=--//
	var t Team   
	_ = json.NewDecoder(r.Body).Decode(&t)

	if t.IsEmpty(){
		json.NewEncoder(w).Encode("Please send some data")
		return;
	} 


	// generate unique id, string
	// append new team to team	
	// rand.Seed(time.Now().UnixNano())
	// rand.Intn(100)
	t.TeamId = strconv.Itoa(rand.Intn(100))
	DBteam=append(DBteam, t);
	json.NewEncoder(w).Encode(t)
	// return;
	
}

func updateOneTeam(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update Team")
	w.Header().Set("Content-Type", "application/json")	
	params :=mux.Vars(r)

	for index, t := range DBteam{
		if t.TeamId == params["id"]{
			// =--=-=-first deletee that value in index position -=--=-=//
			DBteam = append(DBteam[:index], DBteam[index+1:]...)
			var t Team   
			_ = json.NewDecoder(r.Body).Decode(&t)
			t.TeamId = params["id"]	
			DBteam = append(DBteam, t)
			json.NewEncoder(w).Encode(t)
			return;

		}
	}
}

func deleteOneTeam(w http.ResponseWriter, r *http.Request){
	fmt.Println("Detete team");
	w.Header().Set("Content-type", "application/json");
	params:=mux.Vars(r);

	//==--==- loop, id, remove(index, index+1) =--==-//
	for index, t :=range DBteam{

		if t.TeamId==params["id"]{
			DBteam = append(DBteam[:index], DBteam[index+1:]...);
			break;
		}

	}

}