package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer test\n")
	writer.Close()
	os.Exit(0)
}
