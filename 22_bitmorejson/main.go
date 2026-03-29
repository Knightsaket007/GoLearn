package main

import (
	"encoding/json"
	"fmt"
)

type team struct {
	Name           string `json:"TeamName"`
	Members        int16  `json:"bande"`
	Rank           int    `json:"position"`
	Password       string
	PopularMembers []string `json:"heros,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJSon()
	DecodeJson()
}

func EncodeJSon() {
	MCU_team := []team{
		{"Avengers", 10, 1, "hydra ki jai", []string{"Iron Man", "spider man", "Captain America"}},
		{"x men", 12, 2, "hello123", []string{"wolverine", "magnito", "X javier"}},
		{"thunderbolt", 5, 3, "marvel123", []string{"Sentry", "winter soldier"}},
		{"teamX", 6, 4, "marvel1234", []string{}},
	}

	finalJson, err := json.MarshalIndent(MCU_team, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
	   {
           "TeamName": "Avengers",
           "bande": 10,
           "position": 1,
           "Password": "",
           "heros": ["Iron Man", "spider man", "Captain America"]
       }
	`)

	var teams team;

	checkValid := json.Valid(jsonDatafromWeb)

	if checkValid{
		fmt.Println("Valid Json");
		json.Unmarshal(jsonDatafromWeb, &teams)
		
		fmt.Printf("%#v\n", teams)
	}else{
		fmt.Println("Json is not valid")
	}


	// -===- some case where u just want to add key value ==----=--==//

	var myonlineData map[string] interface{}

	json.Unmarshal(jsonDatafromWeb, &myonlineData);
	fmt.Printf("%#v\n", myonlineData);

	for k,v :=range myonlineData{
		fmt.Printf("key is %v and value is %v and type is %T", k, v,v)
	}

}
