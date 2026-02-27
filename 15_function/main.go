package main

import "fmt"

func main() {
	fmt.Println("This is function");

	// sum:= adder(34, 32);
	// fmt.Printf("sum of 2 number is... %v", sum)
	sum, length:=proAdder(1,2,3,4)
	fmt.Printf("sum of multiple numbers is: %v and length of input: %v", sum, length)
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo;
}

func proAdder(values ...int) (int, int) {
	sum:=0
	if len(values)<=0{
		return sum, len(values);
	}

	for _, val := range values{
		// fmt.Println("val...", val)
		sum+=val;
	}
	return sum, len(values);
}