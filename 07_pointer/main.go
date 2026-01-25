package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer");

	mynumber:=100;
	var numberPointer = &mynumber;

	fmt.Println("the value of numberPointer address", numberPointer);
	fmt.Println("the value of numberPointer value", *numberPointer);

	// numberPointer=33
	*numberPointer=*numberPointer + 100;

	fmt.Println("this is value of mynumber...", mynumber)


}