package main

import "fmt"

func main(){
	fmt.Println("structs in golang")

	me:=User{"Saket", "saket@mail.com", true, 25}

	fmt.Println(me);
	fmt.Printf("Myy details are: %+v\n", me)
	fmt.Printf("My name is %v, and my age is %v \n", me.Name, me.Age)


	// -==-=-= Use of Method --=-=-//
	 me.GetStaus()
	 me.GetStaus()

}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}


// =--=-=-= create a method =-=-=-=-//

func (u User) GetStaus(){
	fmt.Println("status is ...", u.Status);
}

func (u User) IncrementAge(){

	// this will not change the origin value. 
	u.Age= u.Age+1

	fmt.Println("Age  incremented ...", u.Age);
}