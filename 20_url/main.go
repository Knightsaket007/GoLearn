package main

import (
	"fmt"
	"net/url"
)

const myurl="https://www.helloworld.com?test=123&payment=321"

func main() {
	fmt.Println("Welcome to url handler")
	fmt.Println(myurl)

	// parsing

	result,_:= url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qparams:= result.Query()
	fmt.Printf("The type of this qparams are : %T\n", qparams)

	fmt.Println(qparams["payment"])

	for _, val:= range qparams{
		fmt.Println("the value of query sting:", val)
	}


	// =--=- Make new URL -=-=-=-//

	partsOfUrl := &url.URL{
		Scheme: "https",
		Path: "localhost:3000",
		RawQuery: "test=111",
	}

	fmt.Println("new url:", partsOfUrl)
}