package main

import "fmt"

func main() {
	fromDollarToEuro := 0.92
	fromDollarToRuble := 84.19
	fromEuroToRuble := fromDollarToRuble / fromDollarToEuro
	fmt.Printf("%.2f\n", fromEuroToRuble)
	userInputSum := userInputSum()
	fmt.Println(userInputSum)
}

func userInputSum() float64 {
	var userInputSum float64
	fmt.Print("Enter the sum to convert: ")
	fmt.Scan(&userInputSum)
	return userInputSum
}
