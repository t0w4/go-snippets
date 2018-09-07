package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub finished")
		time.Sleep(3 * time.Second)
		cancel()
		time.Sleep(100 * time.Second)
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
