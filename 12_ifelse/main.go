package main

import "fmt"

func main() {
	fmt.Printf("This is if else");

	loginCount:=23;

	// result:="";
	var result string;

	if loginCount < 10{
		result="Regular user"
	} else if loginCount > 10{
		result="Watch out"
	}else {
		result="exact 10 login count"
	}

	fmt.Printf("Result is %s",result);

}