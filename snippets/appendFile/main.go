package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		os.Exit(1)
	}
	_, err = file.Write([]byte("row1\n"))
	if err != nil {
		fmt.Printf("write file error: %v\n", err)
		os.Exit(1)
	}
	err = append(fileName)
	if err != nil {
		fmt.Printf("append file error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func append(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, "row2\n")
	if err != nil {
		return err
	}
	return nil
}
