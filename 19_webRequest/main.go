package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

    defer response.Body.Close() //=-=--=-close the connection =-=-=-=-//

	databyte, err := io.ReadAll(response.Body);

	if err !=nil{
		panic(err);
	}

	content:= string(databyte)
	fmt.Println("This Api content...", content)
}
