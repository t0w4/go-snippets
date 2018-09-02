package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "name: %s, age: %d, rate; %.2f", "bob", 18, 0.8)
}
