package main

import "fmt"

const fromDollarToEuro = 0.92
const fromDollarToRuble = 84.19
const fromEuroToRuble = fromDollarToRuble / fromDollarToEuro

func main() {
	fmt.Printf("%.2f\n", fromEuroToRuble)
}

func userInput() (float64, string, string) {
	var userInputSum float64
	var userInputOriginalCurrency string
	var userInputTargetCurrency string
	fmt.Print("Enter the sum to conver: ")
	fmt.Scan(&userInputSum)
	fmt.Print("Enter the original currency to convert(E/D/R): ")
	fmt.Scan(&userInputOriginalCurrency)
	fmt.Print("Enter the target currency to convert(E/D/R): ")
	fmt.Scan(&userInputTargetCurrency)
	return userInputSum, userInputOriginalCurrency, userInputTargetCurrency

}

func calculationFuncion(userInputSum float64, userInputOriginalCurrency string, userInputTargetCurrency string) {
	if userInputOriginalCurrency == "E" || userInputTargetCurrency == "D" {
		// euroToDollar := userInputSum / fromDollarToEuro
	}
}
