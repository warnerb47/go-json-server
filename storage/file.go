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
	s := &FileStorage{
		Path: path,
	}
	return s
}

func (s *FileStorage) Setpath(p string) {
	if p != "" {
		s.Path = p
	} else {
		s.Path = "./data.json"
	}
}

func (s *FileStorage) Load() map[string]any {
	jsonFile, err := os.Open(s.Path)
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
	err = ioutil.WriteFile(s.Path, byte, 0644)
	if err != nil {
		return err
	}
	return nil
}
