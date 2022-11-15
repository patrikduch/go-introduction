package main

import (
	"fmt"
	"time"
)

func printIt(it string) {
	fmt.Println(it)
}

func main() {
	printIt("Initial print")
	for i := 0; i < 100; i++ {
		go println("Print in loop")
	}

	go func(msg string) {
		fmt.Println(msg)
	}("This is a message using a goroutine")

	time.Sleep(time.Second)

	fmt.Println("End")
}
