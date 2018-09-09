package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "name: %s, age: %d, rate; %.2f\n", "bob", 18, 0.8)
}
