package main

import "fmt"

// Create a new type of cryptoassets
// which is  a slice of strings
type CryptoAssets []string;

func (c CryptoAssets) print() {
	for i, crypto := range c {
		fmt.Println(i, crypto);
	}
}