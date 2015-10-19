//+build ignore
// Simple example that will nuke the contents of a s3 bucket
package main

import (
	"fmt"
	"runtime"
	s3 "s3Manager"
)

func init() {
	// multicorebitches
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func printKey(kb s3.KeyBucket) {
	fmt.Println(kb.Key.Key)
	return
}

func main() {
	c := s3.Connection{}
	err := c.Connect()
	if err != nil {
		panic(err)
	}
	b, err := s3.Getbucket("eta-events-msgpack", &c)
	if err != nil {
		panic(err)
	}
	s3.ApplyToMultiList(b, "", "", printKey)
	//fmt.Printf("waiting for 5 sec..")
	//time.Sleep(5 * time.Second)
}
