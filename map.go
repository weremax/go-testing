package main

import "fmt"

func main() {
	// Reference Type not like Struct which are Value Type
	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	updateMap(colors)

	printMap(colors)
}

func updateMap(c map[string]string) {
	c["white"] = "#ffffff"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}