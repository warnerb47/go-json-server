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

func LoadJson() map[string]interface{} {
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
	fmt.Println(result)
	return result
}

func LoadUsers() Users {
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
