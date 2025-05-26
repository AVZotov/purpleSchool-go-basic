package main

import (
	"encoding/json"
	"fmt"
	"json_sli/bin"
	"json_sli/files"
	"json_sli/storage"
)

func toBytes(binList bins.BinList) ([]byte, error) {
	//data, err := json.MarshalIndent(&binList, "", "") //spaces between ":" separator added by Marshall formatter to pass json validation function.
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

func loadFromFile(filename string) ([]byte, error) {
	data, err := files.LoadFile(filename)
	if err == nil {
		fmt.Printf("Data successfully loaded from: %s\n", filename)
		return data, nil
	}
	fmt.Printf("error loading file: %s\n", err)

	err = storage.CreateEmpty()
	if err != nil {
		fmt.Printf("error creating empty file: %s\n", err)
		return nil, err
	}

	fmt.Println("Empty storage successfully created")
	data, err = storage.Load()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func main() {
	var binList = bins.NewBinList()
	data, err := loadFromFile("someFile.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	binList, err = toBinList(data)
	if err != nil {
		fmt.Println(err)
	}

	binList.Add(bins.NewBin(true, "Nikolay"))
	binList.Add(bins.NewBin(false, "Sergey"))
	binList.Add(bins.NewBin(true, "Andrey"))
	binList.Add(bins.NewBin(false, "Vladimir"))

	data, err = toBytes(*binList)
	err = storage.Save(data)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
