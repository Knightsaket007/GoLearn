package main

import (
	"encoding/json"
	"fmt"
)

type team struct{
	Name string `json:"TeamName"`
	Members int16 `json:"bande"`
	Rank int `json:"position"`
	Password string 
	PopularMembers []string `json:"heros,omitempty"`
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
		{"teamX", 6, 4, "marvel1234", []string{}},
	}

	finalJson, err:= json.MarshalIndent(MCU_team, "", "\t")

	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)


}