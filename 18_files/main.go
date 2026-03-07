package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang");
	conent:="This needs  to go  in file (This is file content)";

	file, err:= os.Create("./myGoFile.txt")

	if err!= nil{
		panic(err)
	}

	length, err:=io.WriteString(file, conent)

	if err != nil{
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()
}