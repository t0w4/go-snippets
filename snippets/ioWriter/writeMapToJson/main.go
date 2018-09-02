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
		os.Exit(1)
	}
	defer file.Close()

	sorce := map[string]string{
		"name":  "bob",
		"state": "tokyo",
	}
	e := json.NewEncoder(file)
	e.SetIndent("", "    ")
	if err := e.Encode(sorce); err != nil {
		fmt.Fprintf(os.Stdout, "json encode err: %v\n", err)
		os.Exit(1)
	}
}
