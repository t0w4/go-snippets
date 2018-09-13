package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	var testData = []byte("0123456789ABCDEFG")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		fmt.Printf("write file error: %v\n", err)
		os.Exit(1)
	}
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("file open error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// map memory
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		fmt.Printf("mmap error: %v\n", err)
		os.Exit(1)
	}
	defer m.Unmap()

	m[9] = 'X'
	m.Flush()

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("read file error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap    : %s\n", m)
	fmt.Printf("file    : %s\n", fileData)
	os.Exit(0)
}
