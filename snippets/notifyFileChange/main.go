package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fsnotify "gopkg.in/fsnotify/fsnotify.v1"
)

func main() {
	signals := make(chan os.Signal, 1)
	// catch SIGINT(Ctrl+C), KILL signal
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("file notify inut error: %v\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event: ", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("create file: %s\n", event.Name)
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("write file: %s\n", event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Printf("remove file: %s\n", event.Name)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Printf("rename file: %s\n", event.Name)
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Printf("chmod file: %s\n", event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Printf("notify error: %v\n", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		fmt.Printf("add watch target error: %v\n", err)
	}
	<-signals
	fmt.Println("watch finish")
	os.Exit(0)
}
