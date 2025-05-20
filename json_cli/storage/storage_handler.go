package storage

import (
	"encoding/json"
	"json_sli/bin"
	"os"
)

const fileName = "bins.json"

func SaveToStorage(binList *bins.BinList) (bool, error) {
	data, err := json.Marshal(&binList)

	if err != nil {
		return false, err
	}

	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		return false, err
	}
	_, err = file.Write(data)

	if err != nil {
		return false, err
	}
	return true, nil
}

func LoadFromStorage(data []byte) (*bins.BinList, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var binList bins.BinList
	err = json.Unmarshal(file, &binList)

	if err != nil {
		return nil, err
	}

	return &binList, nil
}
