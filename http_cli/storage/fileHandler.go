package storage

import (
	"os"
)

type FileHandler struct {
	Filename string
}

func NewFileHandler(filename string) *FileHandler {
	return &FileHandler{
		Filename: filename,
	}
}

func (fh *FileHandler) LoadFile() ([]byte, error) {
	data, err := os.ReadFile(fh.Filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fh *FileHandler) SaveFile(data []byte) error {
	file, err := os.Create(fh.Filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
