package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("Hello World!\n"))
	buffer.WriteString("Write String")
	fmt.Println(buffer.String())
	os.Exit(0)
}
