package main

import "fmt";

type person struct {
	firstName string
	lastName string
}

func main () {
	alex := person{firstName: "Patrik", lastName: "Duch"};
	
	fmt.Println(alex.firstName);
}