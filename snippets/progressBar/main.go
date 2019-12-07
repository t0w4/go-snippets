package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	BLUE = "\033[34m"
	END  = "\033[0m"
)

func main() {
	max := 10
	fmt.Fprint(os.Stdout, BLUE)
	for i := 1; i <= max; i++ {
		fmt.Printf("%s [%d/%d]", strings.Repeat("=", i), i, max)
		time.Sleep(500 * time.Millisecond)
		fmt.Fprintln(os.Stdout, "")
		if i != max {
			fmt.Fprint(os.Stdout, "\033[1A")
		}
	}
	fmt.Println("COMPLETE!")
	fmt.Fprint(os.Stdout, END)
}
