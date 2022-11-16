package main

import "fmt"

// package scope variable declaration
var name string

func main() {

	// short variable declaration and initialization (block scope)
	numbers := []int{1, 2}

	// variable assignment
	name = "Patrik"

	fmt.Printf("Get type of numbers array: %T", numbers)
	fmt.Println(numbers)
	fmt.Println("My name is:", name)
}
