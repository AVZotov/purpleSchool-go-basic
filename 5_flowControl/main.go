package main

import (
	"errors"
	"fmt"
)

func userMenu() {
	var initCurrency, targetCurrency string
	var amount float64
	var err error
	fmt.Println("Enter your currency (accepted currencies: USD, EUR, RUB)")
	for {
		initCurrency, err = getCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	fmt.Println("Enter targeted currency (accepted currencies: USD, EUR, RUB)")
	for {
		targetCurrency, err = getCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if targetCurrency == initCurrency {
			fmt.Println("Currencies matching. Nothing to convert, please try again.")
			continue
		}
		break
	}
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

func getCurrency() (string, error) {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return "", err
	}
	if userInput != "USD" && userInput != "EUR" && userInput != "RUB" {
		return "", errors.New("currency is not supported")
	}
	return userInput, nil
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
		return 0.88
	case "EURUSD":
		return 1.13
	case "USDRUB":
		return 81.49
	case "RUBUSD":
		return 0.012271
	case "EURRUB":
		return 92.84
	case "RUBEUR":
		return 0.010772
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
