package main

import (
	"fmt"
)

func main() {
	var colors = map[string]string{
		"red":   "#ff0000",
		"white": "#fff",
	}

	colors["green"] = "#32sdds"

	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
