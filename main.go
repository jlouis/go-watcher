package main

import (
	"flag"
	"fmt"
	"matilde/watcher"
	"os"
	"os/signal"
)

//Close controls how to close matilde.
// it intercepts ^c and instead of sending SIGTERM
// it makes all the function to return as if they were done with their job.
func Close(done chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				fmt.Printf("Captured %v, saving to disk and exiting..\n", sig)
				close(done)
			}
		}
	}()
}

func main() {
	done := make(chan struct{})
	Close(done)
	path := flag.String("path", "", "location of the  serialized data. This is the directory that will be watched")
	flag.Parse()
	fmt.Printf("--------> starting watcher \n")
	// configure action
	var action watcher.Action
	action.Do = watcher.S3Example
	// start the watcher
	watcher := watcher.Watch
	fmt.Printf("--------> now watching\n")
	watcher(done, *path, action, 100, "msgpack", "xz")
}
