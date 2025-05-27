package files

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func LoadFile(filename string) ([]byte, error) {
	if !isJsonExtension(filename) {
		return nil, errors.New(filename + "not a json file extension")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if !isValidJson(data) {
		return nil, errors.New(filename + "not a valid json file")
	}
	return data, nil
}

func isJsonExtension(fileName string) bool {
	return filepath.Ext(fileName) == ".json"
}

func isValidJson(data []byte) bool {
	return json.Valid(data)
}
