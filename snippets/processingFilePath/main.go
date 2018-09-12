package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// get current dir
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("get working directory error: %v\n", err)
	}
	// join directory and filename
	fullPath := filepath.Join(dir, "main.go")
	fmt.Printf("this file is %s\n", fullPath)

	//split fullPath to dir, file
	fmt.Println(filepath.Split(fullPath))

	// get base
	fmt.Printf("file base is %s\n", filepath.Base(fullPath))

	// get dir
	fmt.Printf("file dir  is %s\n", filepath.Dir(fullPath))

	// get ext
	fmt.Printf("file ext  is %s\n", filepath.Ext(fullPath))

	// separate paths
	fmt.Printf("split paths is %v\n", filepath.SplitList("/var:./:/etc"))

	// clean path
	fmt.Printf("clean path is %s\n", filepath.Clean("../processingFilePath/./../processingFilePath/main.go"))

	// expand env
	fmt.Printf("go src is in %s\n", os.ExpandEnv("${GOPATH}/src"))
}
