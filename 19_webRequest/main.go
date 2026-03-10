package main

import (
	"fmt"
	"net/http"
)

const url ="https://google.com"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url);

	if err!= nil{
		panic(err);
	}

	fmt.Printf("Response is of type: %T\n", response)
}