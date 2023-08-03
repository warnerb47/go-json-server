package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var path string = "./data.json"

func Setpath(p string) {
	if p != "" {
		path = p
	} else {
		path = "./data.json"
	}
}

func LoadJson() map[string]any {
	jsonFile, err := os.Open(path)
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

func WriteJson(key string, data any) error {
	jsonData := LoadJson()
	jsonData[key] = data

	byte, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, byte, 0644)
	if err != nil {
		return err
	}
	return nil
}
