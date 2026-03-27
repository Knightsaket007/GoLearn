package main

import (
	"encoding/json"
	"fmt"
)

type team struct{
	Name string
	Members int16
	rank int
	password string
	popularMembers []string
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJSon()
}

func EncodeJSon(){
	MCU_team:= []team{
		{"Avengers", 10, 1, "hydra ki jai", []string{"Iron Man", "spider man", "Captain America"}},
		{"x men", 12, 2, "hello123", []string{"wolverine", "magnito", "X javier"}},
		{"thunderbolt", 5, 3, "marvel123", []string{"Sentry", "winter soldier"}},
	}

	finalJson, err:= json.Marshal(MCU_team)

	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}