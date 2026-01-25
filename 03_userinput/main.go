package main

import (
	"bufio"
	"fmt"
	"os"
)

var token="fdnfnsns vknfrerer"

func main() {
	fmt.Println("hello world")
	fmt.Println("token vaue is:", token)
	reader:=bufio.NewReader(os.Stdin);
	fmt.Println("Please enter number")
	 input, _:=reader.ReadString('\n');
	 fmt.Println("this is your input", input)
	// fmt=
}