package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error detected")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name(), "/ - (Directory)")
		} else {
			fmt.Println(entry.Name(), " - (File)")
		}
	}
}