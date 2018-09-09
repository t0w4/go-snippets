package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Dial error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "Hello World!\n"
	conn.Write([]byte(message))
}
