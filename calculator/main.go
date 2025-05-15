package main

import (
	"fmt"
)

const (
	USDEUR = 0.88
	USDRUB = 81.49
	EURRUB = 92.84
)
const (
	EURUSD = 1 / USDEUR
	RUBUSD = 1 / USDRUB
	RUBEUR = 1 / EURRUB
)

func userMenu() {
	initCurrency, targetCurrency := getCurrencies()
	var amount float64
	var err error

	fmt.Println("Enter the amount of money to convert:")
	for {
		amount, err = getAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	fmt.Printf("Converting: %.2f%s is %.2f%s\n",
		amount, initCurrency, calculate(initCurrency, targetCurrency, amount), targetCurrency)
}

func getCurrencies() (string, string) {
	var initCurrency, targetCurrency string
	var err error
	fmt.Println("Enter your currency (accepted currencies: USD, EUR, RUB)")
	for {
		initCurrency, err = scanUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !checkCurrency(initCurrency) {
			continue
		}
		break
	}
	fmt.Println("Enter targeted currency (accepted currencies: USD, EUR, RUB)")
	for {
		targetCurrency, err = scanUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !checkCurrency(targetCurrency) {
			continue
		}
		if targetCurrency == initCurrency {
			fmt.Println("Currencies matching. Nothing to convert, please try again.")
			continue
		}
		break
	}
	return initCurrency, targetCurrency
}

func scanUserInput() (string, error) {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return "", err
	}
	return userInput, nil
}

func checkCurrency(currency string) bool {
	if currency != "USD" && currency != "EUR" && currency != "RUB" {
		fmt.Println("Currency is not supported")
		return false
	}
	return true
}

func getAmount() (float64, error) {
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func getFX(initCurrency, targetCurrency string) float64 {
	fx := initCurrency + targetCurrency

	switch fx {
	case "USDEUR":
		return USDEUR
	case "USDRUB":
		return USDRUB
	case "EURRUB":
		return EURRUB
	case "RUBUSD":
		return RUBUSD
	case "RUBEUR":
		return RUBEUR
	case "EURUSD":
		return EURUSD
	default:
		return 0
	}
}

func calculate(initCurrency, targetCurrency string, amount float64) float64 {
	return getFX(initCurrency, targetCurrency) * amount
}

func main() {
	userMenu()
}
