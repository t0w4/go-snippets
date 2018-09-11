package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Listen handle error: %v\n", err)
		os.Exit(1)
	}
	defer li.Close()

	// repeat Accept
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Accept handle error : %v\n", err)
			break
		}
		go func() {
			buffer, err := ioutil.ReadAll(conn)
			if err != nil {
				fmt.Fprintf(os.Stdout, "read request err: %v\n", err)
				conn.Close()
				os.Exit(1)
			}
			fmt.Println(string(buffer))
			conn.Close()
		}()
	}
	os.Exit(0)
}
