package main

import "fmt"

func main() {
	// slice
	cryptoAssets := []string {"polygon", "ethereum"};
	cryptoAssets = addCrypto(cryptoAssets, "dogecoin");

	for i, crypto := range cryptoAssets {
		fmt.Println(i, crypto);
	}
} 

func addCrypto(cryptoAssets []string, newCrypto string) []string {
	cryptoAssets = append(cryptoAssets, newCrypto);
	return cryptoAssets;
}

func testMethod() string {
	return "buy-btc";
}