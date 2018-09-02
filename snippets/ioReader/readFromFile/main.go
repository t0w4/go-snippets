package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(buffer))
}
