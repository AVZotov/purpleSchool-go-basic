package main

import (
	"fmt"
	"json_sli/bin"
	"json_sli/storage"
)

func main() {
	var binList = bin.NewVault(storage.NewJsonDb("bins.json"), "BinList")
	fmt.Println(binList)

	binList.Add(bin.NewBin(true, "Nikolay"))
	binList.Add(bin.NewBin(false, "Sergey"))
	binList.Add(bin.NewBin(true, "Andrey"))
	binList.Add(bin.NewBin(false, "Vladimir"))
	err := binList.Save()
	if err != nil {
		fmt.Println(err)
	}
}
