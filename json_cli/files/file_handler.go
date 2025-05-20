package files

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	if !isJson(filename) {
		return nil, errors.New(filename + " is not a json")
	}

	if !isValidJson([]byte(filename)) {
		return nil, errors.New(filename + " not valid json file format")
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

func isValidJson(filename []byte) bool {
	return json.Valid(filename)
}
