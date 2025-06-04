package bin

import (
	"encoding/json"
	"time"
)

type DataBase interface {
	Write(data []byte) error
	Read() ([]byte, error)
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
	return &VaultWithDb{
		Vault: Vault{
			Name:      vaultName,
			Bins:      []Bin{},
			CreatedAt: time.Now(),
		},
		dataBase: db,
	}
}

func LoadVault(db DataBase) (*VaultWithDb, error) {
	file, err := db.Read()
	if err != nil {
		return nil, err
	}
	var vault *Vault
	vault, err = toVault(file)
	if err != nil {
		return nil, err
	}

	return &VaultWithDb{
		Vault:    *vault,
		dataBase: db,
	}, nil
}

func (v *VaultWithDb) Write() error {
	data, err := v.Vault.toBytes()
	if err != nil {
		return err
	}
	err = v.dataBase.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (v *VaultWithDb) Add(newBin *Bin) {
	v.Vault.Bins = append(v.Vault.Bins, *newBin)
}

func (v *Vault) toBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func toVault(data []byte) (*Vault, error) {
	var vault Vault
	err := json.Unmarshal(data, &vault)
	if err != nil {
		return nil, err
	}
	return &vault, nil
}
