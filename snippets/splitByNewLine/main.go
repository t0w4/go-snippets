package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// ex1
	file1, err := os.Open("test1.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err : %v", err)
		os.Exit(1)
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	// ex2
	file2, err := os.Open("test2.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err : %v", err)
		os.Exit(1)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

}
