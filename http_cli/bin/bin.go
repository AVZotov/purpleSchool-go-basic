package bin

import (
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
		Id:        "",
		Private:   isPrivate,
		CreatedAt: time.Now(),
		Name:      binName,
	}
}
