package main

import "fmt"

func main() {
	const USDToEUR = 0.88
	const USDToRUB = 81.49
	const EURToRUB = USDToRUB / USDToEUR
	fmt.Println(USDToEUR, USDToRUB, EURToRUB)
}
