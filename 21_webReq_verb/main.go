package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetReq()
}

func PerformGetReq(){

	const myurl="http://localhost:8000/get"

	response, err:=http.Get(myurl)

	if err!=nil{
		panic(err);
	}

	defer response.Body.Close();

	fmt.Println("Status Code : ", response.StatusCode);
	fmt.Println("Content length is:", response.ContentLength)

	//--=-=-=-=-  Strings Builder Method =-=-=-=-///
	var responseString strings.Builder
	content, _:= io.ReadAll(response.Body)
	byteCount, _:= responseString.Write(content)
	fmt.Println("Content...", string(content))
	fmt.Println("byteCount...", byteCount)

	//--=-=-=-=-  Simple Method =-=-=-=-///
	// content,_:= io.ReadAll(response.Body)
	// fmt.Println("Content...", string(content))
}