package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
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
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				// set timeout
				conn.SetDeadline(time.Now().Add(5 * time.Second))
				// read request
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// Timeout or Socket Close Check
					netErr, ok := err.(net.Error)
					if ok && netErr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					fmt.Fprintf(os.Stdout, "read request err: %v\n", err)
					os.Exit(1)
				}
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					fmt.Fprintf(os.Stdout, "dump request err: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(string(dump))
				content := "Hello World!"
				// write response (HTTP/1.1)
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
