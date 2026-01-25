package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go World")
	
	reader :=bufio.NewReader(os.Stdin);
	fmt.Println("please enter number between 1 to 5")
	input, _:=reader.ReadString('\n');
	fmt.Println("This is your input", input)
	fmt.Printf("This is your input Type %T \n", input)

	numInput, err:=strconv.ParseFloat(strings.TrimSpace(input),64);

	if err != nil{
		 fmt.Println(err);
		//  return;
		// panic(err)
	}else{
		fmt.Println("sum of same number ", numInput + numInput)
	}

}