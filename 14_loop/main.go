package main

import "fmt"

func main() {
	fmt.Printf("This is loop");

	// days:=[]string {"Sunday", "Monday", "Tuesday","Wednesday", "Thrusday", "Friday", "Saturday"}

	// for i:=0; i< len(days); i++{
	// 	// println("day",i, days[i])
	// 	fmt.Println("day ",i, days[i])
	// }

	// for i:= range days{
	// 	fmt.Println("i", days[i])
	// }

	// for index, val:= range days{
	// 	fmt.Println("i", val, index)
	// }

	// for _, val:= range days{
	// 	fmt.Println("i", val)
	// }


	// =-=--========================================//
	// =-=--========================================//
	
	rougeValue:=1;
	for rougeValue <10{

		if rougeValue == 5 {
			goto lco
		}

		fmt.Println("Value is:", rougeValue);
		rougeValue++;
	}

	lco: fmt.Printf(" i am at 5 ")
}