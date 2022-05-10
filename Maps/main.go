package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Assigning int to map
	// colors[10] = "#ffffff"

	// deleting function from map
	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
