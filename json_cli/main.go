package main

import (
	"encoding/json"
	"fmt"
	"json_sli/bin"
	"json_sli/files"
	"json_sli/storage"
	"os"
)

func toBytes(binList *bins.BinList) ([]byte, error) {
	data, err := json.Marshal(&binList)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func toBinList(data []byte) (*bins.BinList, error) {
	var binList bins.BinList
	err := json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}

func loadFromFile(filename string) []byte {
	data, err := files.LoadFile(filename)
	if err != nil {
		fmt.Println(err)
		err = storage.CreateEmpty()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Empty storage successfully created")
		data, err = storage.Load()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return data
	}
	fmt.Printf("Data successfully loaded from: %s\n", filename)
	return data
}

func main() {

	data := loadFromFile("SomeFile.json")
	binList, err := toBinList(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	binList.Add(bins.NewBin(true, "Nikolay"))
	binList.Add(bins.NewBin(false, "Sergey"))
	binList.Add(bins.NewBin(true, "Andrey"))
	binList.Add(bins.NewBin(false, "Vladimir"))

	data, err = toBytes(binList)
	err = storage.Save(data)
	if err != nil {
		return
	}
}
