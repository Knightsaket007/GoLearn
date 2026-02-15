package main

import "fmt"

func main(){
	fmt.Println("structs in golang")

	me:=User{"Saket", "saket@mail.com", true, 25}

	fmt.Println(me);
	fmt.Printf("Myy details are: %+v\n", me)
	fmt.Printf("My name is %v, and my age is %v", me.Name, me.Age)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}