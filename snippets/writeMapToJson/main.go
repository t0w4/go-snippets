package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.json")
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
	}

	sorce := map[string]string{
		"name":  "bob",
		"state": "tokyo",
	}
	e := json.NewEncoder(file)
	e.SetIndent("", "    ")
	if err := e.Encode(sorce); err != nil {
		fmt.Fprintf(os.Stdout, "json encode err: %v\n", err)
	}
}
