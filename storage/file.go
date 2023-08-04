package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Path string

type FileStorage struct {
	Path string
}

func NewFileStorage(path string) *FileStorage {
	Setpath(path)
	return &FileStorage{}
}

func Setpath(p string) {
	if p != "" {
		Path = p
	} else {
		Path = "./data.json"
	}
}

func (s *FileStorage) Load() map[string]any {
	jsonFile, err := os.Open(Path)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]any
	json.Unmarshal(byteValue, &result)
	return result
}

func (s *FileStorage) Write(key string, data any) error {
	jsonData := s.Load()
	jsonData[key] = data

	byte, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Path, byte, 0644)
	if err != nil {
		return err
	}
	return nil
}
