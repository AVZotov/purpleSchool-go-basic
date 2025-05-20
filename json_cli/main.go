package main

import (
	"encoding/json"
	"fmt"
	"json_sli/bin"
	"json_sli/files"
	"json_sli/storage"
	"os"
)

func LoadFromFile(filename string) bins.BinList {
	var binList bins.BinList
	data, err := files.LoadFile(filename)
	if err != nil {
		fmt.Println(err)
		err = storage.CreateEmptyStorage()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Empty storage successfully created")
		binList, err = storage.LoadFromStorage()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return binList
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Data successfully loaded from: %s\n", filename)
	return binList
}

func main() {

	binList := LoadFromFile("someFile.json")

	binList.Add(bins.NewBin(true, "Alex"))
	binList.Add(bins.NewBin(false, "Sergey"))
	binList.Add(bins.NewBin(true, "Andrey"))
	binList.Add(bins.NewBin(false, "Vladimir"))

	err := storage.SaveToStorage(&binList)
	if err != nil {
		return
	}
}
