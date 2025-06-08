package storage

import (
	"errors"
)

type Encrypter interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type EncryptedJsonDb struct {
	FileHandler FileHandler
	Encrypter   Encrypter
}

func NewEncryptedJsonDb(fh FileHandler, enc Encrypter) *EncryptedJsonDb {
	return &EncryptedJsonDb{
		FileHandler: fh,
		Encrypter:   enc,
	}
}

func (db *EncryptedJsonDb) Write(data []byte) error {
	encryptedData, err := db.Encrypter.Encrypt(data)
	if err != nil {
		return err
	}
	err = db.FileHandler.SaveFile(encryptedData)
	if err != nil {
		return err
	}
	return nil
}

func (db *EncryptedJsonDb) Read() ([]byte, error) {
	if !isJsonExtension(db.FileHandler.GetFilename()) {
		return nil, errors.New("wrong filename: not a json extension")
	}
	data, err := db.FileHandler.LoadFile()
	if err != nil {
		return nil, err
	}
	decryptedData, err := db.Encrypter.Decrypt(data)
	if !isValidJson(&decryptedData) {
		return nil, errors.New("invalid json data")
	}
	return decryptedData, nil
}
