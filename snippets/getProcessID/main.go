package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())  // get process id
	fmt.Println(os.Getppid()) // get parent process id
}
