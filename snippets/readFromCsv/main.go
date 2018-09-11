package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err : %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	cr := csv.NewReader(r)
	for {
		l, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stdout, "read csv err : %v\n", err)
			os.Exit(1)
		}
		fmt.Println(l[2], l[6:9])
	}
	os.Exit(0)
}
