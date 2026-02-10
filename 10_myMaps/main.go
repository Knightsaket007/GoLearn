package main

import "fmt"

func main() {

	fmt.Print("Map in golang")

	language:=make(map[string]string)

	language["go"]="golang";
	language["js"]="Javascript";
	language["ts"]="Typescript";

	fmt.Println("List of langusge...", language)
}