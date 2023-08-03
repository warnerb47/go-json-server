package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Path string = "./data.json"

func LoadJson() map[string]any {
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

func WriteJson(key string, data any) error {
	jsonData := LoadJson()
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
