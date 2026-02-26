package main

import "fmt"

func main() {

	fmt.Print("Map in golang")

	language:=make(map[string]string)

	language["go"]="golang";
	language["js"]="Javascript";
	language["rb"]="Ruby";
	language["ts"]="Typescript";

	fmt.Println("List of langusge...", language)
	fmt.Println("Js short of...", language["js"])
	
	// =-=--= Delete -=-=-=//
	delete(language,"rb")
	fmt.Println("List of langusge...", language)

	// --=--= loop =-=-=-//
	for key ,value := range language{
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}
}