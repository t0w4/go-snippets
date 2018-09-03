package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
