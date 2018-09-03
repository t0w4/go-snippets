package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Fprintf(os.Stdout, "connect server error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	// ex1
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	io.Copy(os.Stdout, conn)

	// ex2
	io.WriteString(conn, "GET / HTTP/1.0\r\n\r\n")
	io.Copy(os.Stdout, conn)

}
