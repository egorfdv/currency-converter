package main

import (
	"errors"
	"fmt"
)

const USDToEUR = 0.88
const USDToRUB = 81.49
const EURToRUB = USDToRUB / USDToEUR

func main() {
	fmt.Println("*** Currency Converter ***")
	for {
		originalCurrency, errOriginalCurrency := getUserInputOriginalCurrency()
		if errOriginalCurrency != nil {
			fmt.Println("Error! Select the correct currency")
			continue
		}
		amountToConvert, errAmountToConvert := getUserInputAmountToConvert()
		if errAmountToConvert != nil {
			fmt.Println("Error! The sum must be greater than 0")
			continue
		}
		targetCurrency, errTargetCurrency := getUserInputTargetCurrency(originalCurrency)
		if errTargetCurrency != nil {
			fmt.Println("Error! Select the correct currency")
			continue
		}
		result := userDataCalculation(originalCurrency, amountToConvert, targetCurrency)
		fmt.Printf("Итог: %.2f\n", result)
		userChoice := userMenu()
		if !userChoice {
			break
		} else {
			continue
		}
	}
}

func getUserInputOriginalCurrency() (string, error) {
	var originalСurrency string
	fmt.Print("Select the original currency(USD/EUR/RUB): ")
	fmt.Scan(&originalСurrency)
	if originalСurrency == "USD" || originalСurrency == "EUR" || originalСurrency == "RUB" {
		return originalСurrency, nil
	}
	return "", errors.New("ERROR")
}

func getUserInputAmountToConvert() (float64, error) {
	var amountToConvert float64
	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&amountToConvert)
	if amountToConvert < 0.0 {
		return 0.0, errors.New("ERROR")
	}
	return amountToConvert, nil
}

func getUserInputTargetCurrency(originalСurrency string) (string, error) {
	var targetCurrency string
	switch originalСurrency {
	case "USD":
		fmt.Print("Select the original currency(EUR/RUB): ")
		fmt.Scan(&targetCurrency)
	case "EUR":
		fmt.Print("Select the original currency(USD/RUB): ")
		fmt.Scan(&targetCurrency)
	case "RUB":
		fmt.Print("Select the original currency(USD/EUR): ")
		fmt.Scan(&targetCurrency)
	}
	if targetCurrency == "USD" || targetCurrency == "EUR" || targetCurrency == "RUB" {
		return targetCurrency, nil
	}
	return "", errors.New("ERROR")
}

func userDataCalculation(originalСurrency string, amountToConvert float64, targetCurrency string) float64 {
	var res float64
	if originalСurrency == "USD" && targetCurrency == "EUR" {
		res = amountToConvert * USDToEUR
	} else if originalСurrency == "USD" && targetCurrency == "RUB" {
		res = amountToConvert * USDToRUB
	} else if originalСurrency == "EUR" && targetCurrency == "USD" {
		res = amountToConvert / USDToEUR
	} else if originalСurrency == "EUR" && targetCurrency == "RUB" {
		res = amountToConvert * EURToRUB
	} else if originalСurrency == "RUB" && targetCurrency == "USD" {
		res = amountToConvert / USDToRUB
	} else if originalСurrency == "RUB" && targetCurrency == "EUR" {
		res = amountToConvert / EURToRUB
	}
	return res
}

func userMenu() bool {
	var userChoice string
	fmt.Print("Do you want to repeat the conversion?(Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
