package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	for {
		b, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stdout, "read line err : %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%#v\n", string(b))
	}
	os.Exit(0)
}
