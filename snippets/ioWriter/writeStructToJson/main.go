package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.json")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
	}
	defer file.Close()

	type Address struct {
		ZipCode  string `json:"zip_code"`
		State    string `json:"state"`
		Address1 string `json:"address_1"`
		Address2 string `json:"address_2"`
	}
	type User struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	address := Address{
		ZipCode:  "111-1111",
		State:    "tokyo",
		Address1: "minatoku",
		Address2: "odaiba",
	}

	user := User{
		Name:    "taro",
		Age:     18,
		Address: address,
	}

	e := json.NewEncoder(file)
	e.SetIndent("", "    ")
	if err := e.Encode(user); err != nil {
		fmt.Fprintf(os.Stdout, "json encode err: %v\n", err)
	}
}
