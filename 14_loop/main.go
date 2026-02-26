package main

import "fmt"

func main() {
	fmt.Printf("This is loop");

	days:=[]string {"Sunday", "Monday", "Tuesday","Wednesday", "Thrusday", "Friday", "Saturday"}

	// for i:=0; i< len(days); i++{
	// 	// println("day",i, days[i])
	// 	fmt.Println("day ",i, days[i])
	// }

	for i:= range days{
		fmt.Println("i", days[i])
	}
}