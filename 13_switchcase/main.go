package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is switch case")

	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open ")
	case 2:
		fmt.Println("Dice value 2 spot ")
	case 3:
		fmt.Println("Dice value 3 spot")
	case 4:
		fmt.Println("Dice value 4 spot")
	case 5:
		fmt.Println("Dice value 5 spot")
	case 6:
		fmt.Println("Dice value 6 spot and roll dice again")
	}
}
