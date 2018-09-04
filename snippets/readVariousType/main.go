package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var line = "123 45.6 8.0e4 test"
	r := strings.NewReader(line)
	var i int
	var f, g float64
	var s string
	_, err := fmt.Fscan(r, &i, &f, &g, &s)
	if err != nil {
		fmt.Fprintf(os.Stdout, "line scan err : %v", err)
		os.Exit(1)
	}
	fmt.Printf("i=%#v, f=%#v, g=%#v, s=%#v\n", i, f, g, s)
}
