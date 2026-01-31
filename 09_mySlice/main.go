package main

import "fmt"

func main() {
	fmt.Printf("Welcome to Slice");

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList... %T\n", fruitList);

	fruitList=append(fruitList[:2]);
	fmt.Println("Fruitlist is...", fruitList)

}