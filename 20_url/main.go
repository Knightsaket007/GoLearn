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
	fmt.Println(qparams)
}