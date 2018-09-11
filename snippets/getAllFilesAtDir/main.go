package main

import (
	"fmt"
	"os"
)

func main() {
	dirName := "./"
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Printf("can't open %s\n", dirName)
		os.Exit(1)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("can't read dir: %v\n", err)
		os.Exit(1)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("Dir: %s %s\n", fileInfo.Mode(), fileInfo.Name())
		} else {
			fmt.Printf("File: %s %s\n", fileInfo.Mode(), fileInfo.Name())
		}
	}
	os.Exit(0)
}
