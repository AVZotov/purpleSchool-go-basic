package storage

import (
	"encoding/json"
	"errors"
	"path/filepath"
)

type FileHandler interface {
	SaveFile(data []byte) error
	LoadFile() ([]byte, error)
	GetFilename() string
}

type JsonDb struct {
	FileHandler FileHandler
}

func NewJsonDb(fh FileHandler) *JsonDb {
	return &JsonDb{
		FileHandler: fh,
	}
}

func (db *JsonDb) Write(data []byte) error {
	err := db.FileHandler.SaveFile(data)
	if err != nil {
		return err
	}
	return nil
}

func (db *JsonDb) Read() ([]byte, error) {
	if !isJsonExtension(db.FileHandler.GetFilename()) {
		return nil, errors.New("wrong filename: not a json extension")
	}
	data, err := db.FileHandler.LoadFile()
	if err != nil {
		return nil, err
	}
	if !isValidJson(&data) {
		return nil, errors.New("invalid json data")
	}
	return data, nil
}

func isJsonExtension(fileName string) bool {
	return filepath.Ext(fileName) == ".json"
}

func isValidJson(data *[]byte) bool {
	return json.Valid(*data)
}
