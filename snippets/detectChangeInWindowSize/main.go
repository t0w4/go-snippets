package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func GetWindowSize(fd int) *Size {
	ws := &Window{}
	var err error
	ws.Row, ws.Column, err = terminal.GetSize(fd)
	if err != nil {
		fmt.Printf("get window sieze error: %v", err)
		os.Exit(1)
	}
	return &ws.Size
}

type Size struct {
	Row    int
	Column int
}

type Window struct {
	Size
}

func main() {
	signalChan := make(chan os.Signal, 1)
	// catch SIGINT(Ctrl+C), KILL signal, and window size changes
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGWINCH,
	)
	ws := GetWindowSize(syscall.Stdin)
	fmt.Printf("Row %d : Column %d\n", ws.Row, ws.Column)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			// SIGINT(Ctrl+C)
			case syscall.SIGINT:
				exitChan <- 130

			// kILL signal
			case syscall.SIGTERM:
				exitChan <- 143

			case syscall.SIGWINCH:
				ws := GetWindowSize(syscall.Stdin)
				fmt.Printf("Row %d : Column %d\n", ws.Row, ws.Column)

			default:
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
