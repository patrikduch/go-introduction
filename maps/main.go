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
}
