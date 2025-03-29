package main

import "fmt"

func main() {
	fromDollarToEuro := 0.92
	fromDollarToRuble := 84.19
	fromEuroToRuble := fromDollarToRuble / fromDollarToEuro
	fmt.Printf("%.2f\n", fromEuroToRuble)
}

func userInput() (float64, string, string) {
	var userInputSum float64
	var userInputOriginalCurrency string
	var userInputTargetCurrency string
	fmt.Print("Enter the sum to conver: ")
	fmt.Scan(&userInputSum)
	fmt.Print("Enter the original currency to convert: ")
	fmt.Scan(&userInputOriginalCurrency)
	fmt.Print("Enter the target currency to convert: ")
	fmt.Scan(&userInputTargetCurrency)
	return userInputSum, userInputOriginalCurrency, userInputTargetCurrency

}

// func calculationFuncion(userInputSum float64) {

// }
