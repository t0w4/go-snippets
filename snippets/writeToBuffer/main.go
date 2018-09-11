package main

import (
	"bytes"
	"os"

	"github.com/golang/go/src/fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("Hello World!\n"))
	buffer.WriteString("Write String")
	fmt.Println(buffer.String())
	os.Exit(0)
}
