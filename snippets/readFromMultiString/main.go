package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	l1 := bytes.NewBufferString("line1\n")
	l2 := bytes.NewBufferString("line2\n")
	l3 := bytes.NewBufferString("line3\n")
	r := io.MultiReader(l1, l2, l3)
	io.Copy(os.Stdout, r)
	os.Exit(0)
}
