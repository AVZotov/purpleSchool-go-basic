package main

import (
	"errors"
	"fmt"
	"json_sli/bin"
	"json_sli/config"
	"json_sli/encrypter"
	"json_sli/files"
	"json_sli/storage"
	"os"
)

func main() {
	filename := "encryptedBins.json"
	fh := files.NewFileHandler(filename)
	cfg := config.NewEnvConfig()
	enc := encrypter.NewEncrypter(cfg)
	db := storage.NewEncryptedJsonDb(fh, enc)

	securedVault, err := bin.LoadVault(db)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			fmt.Printf("%s does not exist, creating new vault\n", filename)
			securedVault = bin.NewVault(db, "securedVault")
		default:
			fmt.Println(err)
			return
		}
	}

	securedVault.Add(bin.NewBin(false, "Alex"))

	fmt.Println(securedVault.Vault)
	err = securedVault.Write()
	if err != nil {
		fmt.Println(err)
	}
}
