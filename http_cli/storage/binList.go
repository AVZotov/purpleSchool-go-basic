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

func NewBinList() *BinListWithDb {
	db := NewLocalDb()
	return &BinListWithDb{
		BinList: BinList{
			Data: []Bin{},
		},
		Db: *db,
	}
}

// Create new record from metadata received and add this data
// to local storage. It loads data from local storage. Converts metadata to Bin
// add Bin to BinList and saves data back to local storage
func (bl *BinListWithDb) Create(metadata []byte) error {
	err := bl.Read()
	if err != nil {
		return err
	}
	bin, err := toBin(metadata)
	if err != nil {
		return err
	}

	bl.add(bin)

	err = bl.saveToLocalStorage()
	if err != nil {
		return err
	}
	return nil
}

// Read all data from local storage and add it to BinList
// A successful call returns err == nil, not err == os.PathError
// Because absence of local storage means first start of program or no data stored yet
func (bl *BinListWithDb) Read() error {
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

// Delete record from local storage by ID. It loads local storage and convert it to BinList
// metadata received converted to Bin, then BinList is searched by ID to remove data from local
// storage. If record successfully removed, local storage resaved with updated data
func (bl *BinListWithDb) Delete(id string) error {
	err := bl.Read()
	if err != nil {
		return err
	}
	if !bl.BinList.deleteById(id) {
		return fmt.Errorf("record with id: %s not found in local storage", id)
	}

	err = bl.saveToLocalStorage()
	if err != nil {
		return err
	}
	return nil
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

func (bl *BinListWithDb) add(bin *Bin) {
	bl.BinList.Data = append(bl.BinList.Data, *bin)
}

func (bl *BinListWithDb) saveToLocalStorage() error {
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
