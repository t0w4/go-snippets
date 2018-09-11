package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	c := 1024
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("read rand bytes err: %v\n", err)
		os.Exit(1)
	}

	f, err := os.Create("tmp.txt")
	if err != nil {
		fmt.Printf("cread file err: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		fmt.Printf("write file err: %v, byte: %d\n", err, n)
		os.Exit(1)
	}
	os.Exit(0)
}
