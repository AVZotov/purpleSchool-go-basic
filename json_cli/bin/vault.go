package bin

import (
	"encoding/json"
	"fmt"
	"time"
)

type DataBase interface {
	Save(data []byte) error
	Load() ([]byte, error)
}

type Vault struct {
	Name      string    `json:"name"`
	Bins      []Bin     `json:"bins"`
	CreatedAt time.Time `json:"created_at"`
}

type VaultWithDb struct {
	Vault    Vault
	dataBase DataBase
}

func NewVault(db DataBase, vaultName string) *VaultWithDb {
	file, err := db.Load()
	if err != nil {
		fmt.Println(err)
		return &VaultWithDb{
			Vault: Vault{
				Name:      vaultName,
				Bins:      []Bin{},
				CreatedAt: time.Now(),
			},
			dataBase: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err)
		return &VaultWithDb{
			Vault: Vault{
				Name:      vaultName,
				Bins:      []Bin{},
				CreatedAt: time.Now(),
			},
			dataBase: db,
		}
	}
	return &VaultWithDb{
		Vault:    vault,
		dataBase: db,
	}
}

func (v *VaultWithDb) Save() error {
	data, err := v.Vault.ToBytes()
	if err != nil {
		return err
	}
	err = v.dataBase.Save(data)
	if err != nil {
		return err
	}
	return nil
}

func (v *VaultWithDb) Add(newBin *Bin) {
	v.Vault.Bins = append(v.Vault.Bins, *newBin)
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return file, nil
}
