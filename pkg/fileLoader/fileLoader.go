package fileLoader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    string `json:"age"`
	Social Social `json:"social"`
}

type Users struct {
	Users []User `json:"users"`
}

func loadJson() map[string]interface{} {
	jsonFile, err := os.Open("./users.json")
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

func loadUsers() Users {
	jsonFile, err := os.Open("./users.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var users Users
	json.Unmarshal(byteValue, &users)
	fmt.Println(users)
	return users
}

func getKeys(jsonData map[string]interface{}) {
	for key := range jsonData {
		fmt.Println(key)
	}
}

func GetData() {
	result := loadJson()
	getKeys(result)
}
