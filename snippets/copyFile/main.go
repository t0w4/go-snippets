package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fo, err := os.Open("old.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err : %v\n", err)
		os.Exit(1)
	}
	defer fo.Close()
	fn, err := os.Create("new.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err : %v\n", err)
		os.Exit(1)
	}
	defer fn.Close()
	s, err := io.Copy(fn, fo)
	if err != nil {
		fmt.Fprintf(os.Stdout, "copy file err : %v, size: %d\n", err, s)
		os.Exit(1)
	}
}
