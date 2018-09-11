package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	// ex2
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stdout, "read file err : %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%#v\n", scanner.Text())
	}

	// ex3
	rd := bufio.NewReader(os.Stdin)
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stdout, "read line err : %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%#v\n", string(line))
	}
	os.Exit(0)
}
