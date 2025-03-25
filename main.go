package main

import "fmt"

func main() {
	fromDollarToEuro := 0.92
	fromDollarToRuble := 84.19
	fromEuroToRuble := fromDollarToRuble / fromDollarToEuro
	fmt.Printf("%.2f", fromEuroToRuble)
}
