package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

type BinList struct {
	Data []Bin
}

type BinListWithDb struct {
	BinList BinList
	Db      LocalDb
}

func NewBinList() (*BinListWithDb, error) {
	db := NewLocalDb()
	bl := &BinListWithDb{
		BinList: BinList{
			Data: []Bin{},
		},
		Db: *db,
	}
	err := bl.syncWithStorage()
	if err != nil {
		return nil, err
	}
	return bl, nil
}

func (bl *BinListWithDb) Create(metadata []byte) error {
	dataBin, err := toBin(metadata)
	if err != nil {
		return err
	}
	bl.add(dataBin)
	bytes, err := toBytes(&bl.BinList)
	if err != nil {
		return err
	}
	err = bl.Db.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (bl *BinListWithDb) Read() ([]byte, error) {
	data, err := bl.Db.Read()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// func (bl *BinListWithDb) Update(d Bin) error {
//return error
//}

func (bl *BinListWithDb) Delete(d Bin) error {
	return nil
	// TODO remove local bin from list by ID and re save local file
}

func (bl *BinList) deleteById(id string) bool {
	for i, bin := range bl.Data {
		if bin.ID == id {
			bl.Data = slices.Delete(bl.Data, i, i+1)
			return true
		}
	}
	return false
}

func (bl *BinListWithDb) syncWithStorage() error {
	bytes, err := bl.Db.Read()
	pathError := &os.PathError{}
	if err != nil {
		if errors.As(err, &pathError) {
			fmt.Println("No local data detected")
			return nil
		}
		return err
	}
	var data *BinList
	data, err = toBinList(bytes)
	if err != nil {
		return err
	}
	bl.BinList = *data
	return nil
}

func (bl *BinListWithDb) add(bin *Bin) {
	bl.BinList.Data = append(bl.BinList.Data, *bin)
}

func toBytes(binList *BinList) ([]byte, error) {
	bytes, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func toBinList(data []byte) (*BinList, error) {
	var bl BinList
	err := json.Unmarshal(data, &bl)
	if err != nil {
		return nil, err
	}
	return &bl, nil
}
