package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User have user's info
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Address  Address   `json:"address"`
	Comments []Comment `json:"comments"`
}

// Address have user's address
type Address struct {
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

// Comment have user's comment
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	f, err := os.Open("user.json")
	if err != nil {
		fmt.Printf("Open JSON File Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var user User
	if err = dec.Decode(&user); err != nil {
		fmt.Printf("JSON Parse Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(user)
}
