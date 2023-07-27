package fileLoader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJson() map[string]any {
	jsonFile, err := os.Open("./data.json")
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
	err = ioutil.WriteFile("data.json", byte, 0644)
	if err != nil {
		return err
	}
	return nil
}
