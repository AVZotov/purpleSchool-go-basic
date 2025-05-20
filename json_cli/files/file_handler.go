package files

import (
	"errors"
	"os"
	"strings"
)

func LoadFile(filename string) ([]byte, error) {
	if !isJson(filename) {
		return nil, errors.New(filename + " is not a json")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func isJson(fileName string) bool {
	formatted := strings.Split(fileName, ".")
	if len(formatted) < 2 {
		return false
	}
	return strings.ToLower(formatted[len(formatted)-1]) == "json"
}
