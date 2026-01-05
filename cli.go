package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("All arguments: ", os.Args)
	fmt.Println("Program Name: ", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("First argument: ", os.Args[1])
	}
}