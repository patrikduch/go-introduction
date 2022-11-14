package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	patrik := person{firstName: "Patrik", lastName: "Duch"}
	
	fmt.Println(patrik.firstName)
}
