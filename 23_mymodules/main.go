package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")

	r:=mux.NewRouter();
	r.HandleFunc("/", serveHome).Methods("GET")

    log.Fatal(http.ListenAndServe(":4000", r))
}

//=--=-==- API -=--==-//

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang...</h1>"))
}