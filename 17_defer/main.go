package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is defer.")

	defer fmt.Println("\n this is first exicuition, but apply on last queue because of defer keyword");
	fmt.Println("this is second execution but run first because there is no defer.")

	mydefer()
	
}

func mydefer(){
	for i:=0 ; i<10 ; i++{
		defer fmt.Print(i)
	}
}


// -== Explaination -=-//
//*** Output will be **** =>  

// this is second execution but run first because there is no defer.
// 9876543210
//  this is first exicuition, but apply on last queue because of defer keyword

// in defer LIFO (Last In First Out) method is used, so the last defer statement will be executed first and the first defer statement will be executed last. In the above code, the loop will defer the print statements for numbers 0 to 9, and they will be executed in reverse order when the function exits.