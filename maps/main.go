package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Put key to the map
	colors["white"] = "#ffffff"

	// Delete key from map
	delete(colors, "red")

	var mapVariable map[string]string
	anotherMap := make(map[string]string)

	fmt.Println(colors)
	fmt.Println(mapVariable)
	fmt.Println(anotherMap)

	// Print values of map though iterations
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for color", key, "is:", value)
	}
}
