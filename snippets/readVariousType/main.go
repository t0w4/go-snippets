package main

import (
	"fmt"
	"strings"
)

func main() {
	var line = "123 45.6 8.0e4 test"
	r := strings.NewReader(line)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(r, &i, &f, &g, &s)
	fmt.Printf("i=%#v, f=%#v, g=%#v, s=%#v\n", i, f, g, s)
}
