package main

import (
	"fmt"
	"json_sli/bin"
	"json_sli/config"
	"json_sli/encrypter"
	"json_sli/files"
	"json_sli/storage"
)

func main() {
	fh := files.NewFileHandler("encryptedBins.json")
	cfg := config.NewEnvConfig()
	enc := encrypter.NewEncrypter(cfg)
	db := storage.NewEncryptedJsonDb(fh, enc)
	securedVault := bin.NewVault(db, "Secured Vault")

	//securedVault.Add(bin.NewBin(true, "Nikolay"))

	fmt.Println(securedVault.Vault)
	//err := securedVault.Write()
	//if err != nil {
	//	fmt.Println(err)
	//}
}
