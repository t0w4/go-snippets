package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("get work directory error %v\n", err)
		os.Exit(1)
	}
	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == "skipTarget" {
					return filepath.SkipDir
				}
				return nil
			}
			fmt.Println(filepath.Join(path, info.Name()))
			return nil
		})
	if err != nil {
		fmt.Printf("walk directory error %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
