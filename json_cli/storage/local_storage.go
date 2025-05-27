package storage

import (
	"encoding/json"
	"json_sli/bin"
	"os"
)

const fileName = "bins.json"

func Save(data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func Load() ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateEmpty() error {
	var binList = bins.NewBinList()
	data, err := json.Marshal(&binList)
	if err != nil {
		return err
	}
	err = Save(data)
	if err != nil {
		return err
	}
	return nil
}
