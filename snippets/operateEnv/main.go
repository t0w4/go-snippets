package main

import (
	"fmt"
	"os"
)

func main() {
	// Get all env
	fmt.Println(os.Environ())

	// Expand env
	fmt.Println(os.ExpandEnv("${GOPATH}/src"))

	// Set env
	fmt.Println(os.Setenv("test_key", "test_value"))

	// Get env
	fmt.Println(os.Getenv("test_key"))

	// Get env
	fmt.Println(os.LookupEnv("test_key"))
}
