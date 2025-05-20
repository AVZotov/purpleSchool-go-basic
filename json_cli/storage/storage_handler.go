package storage

import (
	"encoding/json"
	"json_sli/bin"
	"os"
)

const fileName = "bins.json"

func SaveToStorage(binList *bins.BinList) error {
	data, err := json.Marshal(&binList)

	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		return err
	}
	_, err = file.Write(data)

	if err != nil {
		return err
	}
	return nil
}

func LoadFromStorage() (bins.BinList, error) {
	file, err := os.ReadFile(fileName)
	var binList bins.BinList

	if err != nil {
		return binList, err
	}

	err = json.Unmarshal(file, &binList)

	if err != nil {
		return binList, err
	}

	return binList, nil
}

func CreateEmptyStorage() error {
	binList := bins.NewBinList()
	err := SaveToStorage(binList)
	if err != nil {
		return err
	}
	return nil
}
