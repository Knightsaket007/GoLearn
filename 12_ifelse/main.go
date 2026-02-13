package main

import "fmt"

func main() {
	fmt.Printf("This is if else");

	loginCount:=23;

	// result:="";
	var result string;
	if loginCount < 10{
		result="Regular user"
	}

	fmt.Printf("Result is %v",result);

}