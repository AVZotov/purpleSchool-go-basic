package bins

import (
	"github.com/google/uuid"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func NewBin(isPrivate bool, binName string) *Bin {
	return &Bin{
		Id:        uuid.NewString(),
		Private:   isPrivate,
		CreatedAt: time.Now(),
		Name:      binName,
	}
}

type BinList struct {
	Bins      []Bin     `json:"bins"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBinList() *BinList {
	return &BinList{
		Bins:      []Bin{},
		CreatedAt: time.Now(),
	}
}

func (b *BinList) Add(newBin *Bin) {
	b.Bins = append(b.Bins, *newBin)
}
