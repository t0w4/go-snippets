package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const BLUE = "\033[34m"
const END = "\033[0m"

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

	makeWindow(ws.Row, ws.Column)

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
				makeWindow(ws.Row, ws.Column)

			default:
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

func makeWindow(row int, column int) {
	print("\033[H\033[2J")
	printFilledLine(row, '+', '-')
	for i := 1; i < (column-1)-1; i++ {
		printFilledLine(row, '|', ' ')
	}
	printFilledLine(row, '+', '-')
}

func printFilledLine(row int, edgeChar byte, fillChar byte) {
	l := make([]byte, row)
	l[0] = edgeChar
	for i := 1; i < (row - 1); i++ {
		l[i] = fillChar
	}
	l[row-1] = edgeChar
	fmt.Println(string(l))
}
