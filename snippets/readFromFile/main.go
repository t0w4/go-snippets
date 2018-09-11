package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	if _, err := file.Read(buf); err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(buf))

	file2, err := os.Open("test2.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer file2.Close()
	buffer, err := ioutil.ReadAll(file2)
	if err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(buffer))
	os.Exit(0)
}
