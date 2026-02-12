package main

import "fmt"

func main(){
	fmt.Println("structs in golang")

	me:=User{"Saket", "saket@mail.com", true, 25}

	fmt.Println(me);
	fmt.Println("Myy details are: %+v", me)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}