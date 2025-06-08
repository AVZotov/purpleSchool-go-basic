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
	selectedOperation := getOperation()
	userSequence := getIntSlice()
	myFunc := map[string]func(*[]int) float64{
		"SUM": calculateSum,
		"AVG": calculateAverage,
		"MED": calculateMedian,
	}
	fmt.Printf("The calculated result is: %.2f",
		myFunc[selectedOperation](&userSequence))
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

func calculateAverage(slice *[]int) float64 {
	var result float64

	for _, value := range *slice {
		result += float64(value)
	}
	return result / float64(len(*slice))
}

func calculateSum(slice *[]int) float64 {
	var result float64
	for _, value := range *slice {
		result += float64(value)
	}
	return result
}

func calculateMedian(slice *[]int) float64 {
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
