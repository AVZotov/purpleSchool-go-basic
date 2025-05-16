package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func menu() {
	userInput := getOperation()
	result := getIntSlice()

	switch userInput {
	case "SUM":
		fmt.Printf("The sum is: %.2f", sum(&result))
	case "AVG":
		fmt.Printf("The avg is: %.2f", avg(&result))
	case "MED":
		fmt.Printf("The mediana is: %.2f", med(&result))
	}
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func getOperation() string {
	var userInput string
	fmt.Print("Please choose the calculator operation" +
		" (SUM, AVG or MED) then press Enter: ")

	for {
		userInput = scan()
		if !checkOperation(userInput) {
			fmt.Print("\nwrong operation please try again: ")
			continue
		}
		return userInput
	}
}

func checkOperation(operation string) bool {
	operationSlice := []string{"SUM", "AVG", "MED"}
	return slices.Contains(operationSlice, operation)
}

func getIntSlice() []int {
	var result []int
	var err error
	fmt.Print("\nEnter integers separated by comma space" +
		"then press Enter: ")

	for {
		result, err = stringConverter(scan())
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			fmt.Println("Please try again:")
			continue
		}
		return result
	}
}

func stringConverter(s string) ([]int, error) {
	formatted := strings.Split(s, ",")
	result := make([]int, 0, len(formatted))

	for _, v := range formatted {
		if strings.TrimSpace(v) != "" {
			intValue, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				return result, err
			}
			result = append(result, intValue)
		}
	}
	return result, nil
}

func avg(slice *[]int) float64 {
	var result float64

	for _, value := range *slice {
		result += float64(value)
	}

	return result / float64(len(*slice))
}

func sum(slice *[]int) float64 {
	var result float64

	for _, value := range *slice {
		result += float64(value)
	}

	return result
}

func med(slice *[]int) float64 {
	sliceCopy := make([]int, len(*slice))
	copy(sliceCopy, *slice)
	slices.Sort(sliceCopy)

	if len(sliceCopy) == 0 {
		return 0
	}

	if len(sliceCopy)%2 == 0 {
		return float64((sliceCopy[len(sliceCopy)/2-1] +
			sliceCopy[len(sliceCopy)/2]) / 2)
	}

	return float64(sliceCopy[len(sliceCopy)/2])
}

func main() {
	menu()
}
