package main

import "fmt"

func main() {
	fmt.Println("This is function");

	// sum:= adder(34, 32);
	// fmt.Printf("sum of 2 number is... %v", sum)
	sum:=proAdder(1,2,3,4)
	fmt.Printf("sum of 2 number is... %v", sum)
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo;
}

func proAdder(values ...int) int {
	sum:=0
	if len(values)<=0{
		return sum;
	}

	for _, val := range values{
		// fmt.Println("val...", val)
		sum+=val;
	}
	return sum;
}