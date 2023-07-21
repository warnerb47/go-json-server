package fileLoader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJson() map[string]interface{} {
	jsonFile, err := os.Open("./data.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	return result
}
