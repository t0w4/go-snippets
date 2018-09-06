package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sun start")
	done := make(chan bool)
	go func() {
		fmt.Println("sub finished")
		time.Sleep(time.Second * 3)
		done <- true
	}()
	<-done
	fmt.Println("all tasks are finished")
}
