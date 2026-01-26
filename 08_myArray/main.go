package main

import "fmt"

func main() {
	fmt.Printf("Welcome to Array");

	var fruitList [4]string

	fruitList[0]="apple"
	fruitList[1]="Banana"
	// fruitList[2]="peach"
	fruitList[3]="mango"

	fmt.Println("Array list..", fruitList)

	fmt.Println("Array list length...", len(fruitList));
}