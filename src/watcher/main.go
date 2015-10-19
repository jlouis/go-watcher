package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

//Close controls how to close watcher.
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
	path := flag.String("path", ".", "Directory to watch for changes. Not recursively.")
	flag.Parse()
	fmt.Printf("--------> starting watcher \n")
	// configure action
	a := XzToS3{}
	a.Connect()
	// start the watcher
	fmt.Printf("--------> now watching: %v\n", *path)
	Watch(done, *path, a, 8, "msgpack", "xz")
}
