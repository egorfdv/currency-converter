package main

import "fmt"

func main() {
	const USDToEUR = 0.88
	const USDToRUB = 81.49
	const EURToRUB = USDToRUB / USDToEUR
	fmt.Println(USDToEUR, USDToRUB, EURToRUB)
}

func getUserInput() string {
	var userChoice string
	fmt.Scan(&userChoice)
	return userChoice
}

func userDataCalculation(num float64, str1 string, str2 string) float64 {
	return num
}
