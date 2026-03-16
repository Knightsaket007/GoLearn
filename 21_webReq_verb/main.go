package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
}

func PerformGetReq(){

	const myurl="https://localhost:8000/get"

	response, err:=http.Get(myurl)

	if err!=nil{
		panic(err);
	}

	defer response.Body.Close();

	fmt.Println("Status Code : ", response.StatusCode);
	fmt.Println("Content length is:", response.ContentLength)

	var responseString strings.Builder
	content,_:= io.ReadAll(response.Body)
}