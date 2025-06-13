package storage

import (
	"encoding/json"
	"fmt"
)

type Bin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewBin(id, name string) Bin {
	return Bin{
		ID:   id,
		Name: name,
	}
}

func (md Bin) String() string {
	return fmt.Sprintf("id: %s\tname: %s\n", md.ID, md.Name)
}

func toBin(metadata []byte) (*Bin, error) {
	var bin Bin
	err := json.Unmarshal(metadata, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, nil
}
