package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// if you want to check SIGTERM, you must build.
func main() {
	signals := make(chan os.Signal, 1)
	// catch SIGINT(Ctrl+C), KILL signal
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("waiting SIGNAL")
	<-signals
	fmt.Println("catch SIGNAL")
	os.Exit(0)
}
