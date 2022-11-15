package main

import "fmt"

func main() {

	messages := make(chan string)
	go test(messages)

	msg := <-messages
	fmt.Println(msg)

	fmt.Println("test")
}

func test(c chan string) chan string {
	c <- "patrik duch"
	return c
}
