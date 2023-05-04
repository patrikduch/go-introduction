package main

import "fmt"

func main() {

	array := []int{1, 2, 3}

	for _, item := range array {
		fmt.Println(item)
	}
}
