package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.Stat(fileName)
	if err == os.ErrNotExist {
		fmt.Printf("%s is not exist\n", fileName)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("file stat error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("-------FIleInfo-------")
	fmt.Printf("file name  : %v\n", file.Name())
	fmt.Printf("file size  : %v\n", file.Size())
	fmt.Printf("change time: %v\n", file.ModTime())
	fmt.Printf("file mode  : %v", file.Mode().String())
	fmt.Printf("(%o)\n", file.Mode().Perm())
	fmt.Printf("directory? : %v\n", file.IsDir())
}
