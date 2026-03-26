package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetReq()
	PerformPostJsonRequest()
	PerformPostFormRequest()
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
	fmt.Println("byteCount...", byteCount)
	fmt.Println("Res...", responseString.String())

	//--=-=-=-=-  Simple Method =-=-=-=-///
	// content,_:= io.ReadAll(response.Body)
	// fmt.Println("Content...", string(content))
}

func PerformPostJsonRequest(){
	const myurl="http://localhost:8000/post"

	// =-=- Fake json payload =-=-//
	requestBody :=strings.NewReader(`
	
	{
	 "Name": "Saket Sourav",
	  "Age": "200",
	  "plateform":"VS Code"
	}
	`);

	response, err:=http.Post(myurl, "application/json", requestBody)

	if err !=nil {
		panic(err)
	}

	defer response.Body.Close();

	content, _:=io.ReadAll(response.Body)

	fmt.Println(string(content))


}

func PerformPostFormRequest(){
	const myurl="http://localhost:8000/postform";

	data:= url.Values{};
	data.Add("firstName", "Saket")
	data.Add("lastName", "Sourav")
	data.Add("email", "saket@go.dev")

	response, err:=http.PostForm(myurl, data);

	if err !=nil{
		panic(err)
	}

	defer response.Body.Close()

	content, _:=io.ReadAll(response.Body)

	fmt.Println("response from Post Form req...", string(content))
}