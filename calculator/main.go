package main

import (
	"fmt"
	"slices"
)

var fxMap = map[string]float64{
	"USDEUR": 0.88,
	"USDRUB": 81.49,
	"EURRUB": 92.84,
	"EURUSD": 1 / 0.88,
	"RUBUSD": 1 / 81.49,
	"RUBEUR": 1 / 92.84,
}

var currencies = []string{"USD", "EUR", "RUB"}

func userMenu() {
	requestedOperation := getOperation()
	var amount float64

	fmt.Print("Enter the amount of money to convert: ")
	amount = getAmount()

	fmt.Printf("Converting: %.2f%s is %.2f%s\n",
		amount, requestedOperation[0:3], calculate(requestedOperation, amount),
		requestedOperation[3:6])
}

func getOperation() string {
	var operation, userInput string
	var err error

	for i := 0; i < 2; i++ {
		fmt.Printf("Enter your %d currency (accepted currencies: %v): ", i+1, currencies)
		for {
			userInput, err = scanUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !checkCurrency(userInput, &currencies) {
				continue
			}
			operation += userInput
			break
		}
	}
	return operation
}

func scanUserInput() (string, error) {
	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return "", err
	}
	return userInput, nil
}

func checkCurrency(currency string, currenciesList *[]string) bool {
	var index int
	if index = slices.Index(*currenciesList, currency); index == -1 {
		fmt.Print("Currency is not supported!\nPlease try again: ")
		return false
	}
	slices.Delete(currencies, index, index+1)
	return true
}

func getAmount() float64 {
	var amount float64
	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	return amount
}

func getFX(operation string, fx *map[string]float64) float64 {
	fxValue, ok := (*fx)[operation]
	if !ok {
		fmt.Printf("operation %s is not supported!", operation)
		return 0
	}
	return fxValue
}

func calculate(operation string, amount float64) float64 {
	return getFX(operation, &fxMap) * amount
}

func main() {
	userMenu()
}
