package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Listen handle error: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is runnninng at localhost:8080")

	// repeat Accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Accept handle error : %v\n", err)
			break
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// read request
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				fmt.Fprintf(os.Stdout, "read request err: %v\n", err)
				conn.Close()
				os.Exit(1)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				fmt.Fprintf(os.Stdout, "dump request err: %v\n", err)
				conn.Close()
				os.Exit(1)
			}
			fmt.Println(string(dump))
			// write response
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World!")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
