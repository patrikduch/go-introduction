package main

func main() {
	// slice
	cryptoAssets := CryptoAssets {"polygon", "ethereum"};
	cryptoAssets = addCrypto(cryptoAssets, "dogecoin");

	cryptoAssets.print();
} 

func addCrypto(cryptoAssets CryptoAssets, newCrypto string) CryptoAssets {
	cryptoAssets = append(cryptoAssets, newCrypto);
	return cryptoAssets;
}

func testMethod() string {
	return "buy-btc";
}