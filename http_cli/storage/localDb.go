package storage

import (
	"encoding/json"
	"errors"
)

const storageName = "metadata.json"

type LocalDb struct {
	Fh *FileHandler
}

func NewLocalDb() *LocalDb {
	fh := NewFileHandler(storageName)
	return &LocalDb{Fh: fh}
}

func (ls *LocalDb) Write(data []byte) error {
	err := ls.Fh.SaveFile(data)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LocalDb) Read() ([]byte, error) {
	data, err := ls.Fh.LoadFile()
	if err != nil {
		return nil, err
	}
	if !isValidJson(&data) {
		return nil, errors.New("invalid json data")
	}
	return data, nil
}

func isValidJson(data *[]byte) bool {
	return json.Valid(*data)
}
