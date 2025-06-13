package main

import (
	"fmt"
	"http_cli/storage"
	"log"
)

func main() {
	ls, err := storage.NewBinList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ls.BinList)
}

//func toBytes(binList *storage.BinList) ([]byte, error) {
//	bytes, err := json.Marshal(binList)
//	if err != nil {
//		return nil, err
//	}
//	return bytes, nil
//}
//
//func main() {
//	listWithDb := storage.NewBinList()
//
//	fmt.Println(listWithDb.BinList.Data)
//	bin3 := storage.NewBin("3", "Chris")
//	bin2 := storage.NewBin("2", "Chris")
//
//	listWithDb.add(&bin2)
//	listWithDb.add(&bin3)
//	fmt.Println(listWithDb.BinList)
//
//	data, _ := toBytes(&listWithDb.BinList)
//	fmt.Println(string(data))
//	listWithDb.Db.Write(data)
//}
