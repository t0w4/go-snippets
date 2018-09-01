package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
	}
	defer file.Close()
	m := io.MultiWriter(file, os.Stdout)
	fmt.Fprintf(m, "name: %s, age: %d, rate; %.2f\n", "bob", 18, 0.8)
}
