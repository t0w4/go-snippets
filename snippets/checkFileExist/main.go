package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	_, err := os.Stat(fileName)
	if err == os.ErrNotExist {
		fmt.Printf("%s is not exist\n", fileName)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("file stat error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s is exist!\n", fileName)
	os.Exit(0)
}
