package main

import (
	"fmt"
	"os"
)

func LoadJson() {
	jsonFile, err := os.Open("users.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
}
