package watcher

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/fsnotify.v1"
	"strings"
	"sync"
)

type Action interface {
	Do(string) error
}

// Watcher watches path  (** not recursively, not supported and not cross platform.**)
// and performs action. Action is a type func(string) that takes the name of changed
// file and does something with it.
// Example defined in this package is an example of am Action.
// this function may panic, if smth goes wrong when:
// starting the watcher, and watching or connecting to s3, since these action happen only once
// and at the beginning it's safer to panic here .
//
// Watcher waits untill all the work is done beefore closing if the done channel is closed.
// This means that anything started before done is closed will finish gratefully.
func Watch(done chan struct{}, path string, action Action, numWorkers int, extensionsWatch string, extensionsExclude string) {
	// init watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Create == fsnotify.Create {
						if strings.Contains(event.Name, extensionsWatch) {
							if !strings.Contains(event.Name, extensionsExclude) {
								action.Do(event.Name)
							}
						}
					}
				case err := <-watcher.Errors:
					log.Errorf("error watching files %v ", err)
				case <-done:
					return
				}
			}
			wg.Done()
		}()
	}
	err = watcher.Add(path)
	if err != nil {
		// NOTE here we keep the panic
		panic(err)
	}
	go func() {
		wg.Wait()
	}()
	<-done
}
