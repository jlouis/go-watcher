// Copyright 2015 Giulio Ungaretti. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package s3Manager  wraps some common s3 actions.
// Credenttials need to be stored in the environment
// See ../example for a simple command line thing .
package s3Manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

// hold wg and counters
var (
	wg          sync.WaitGroup
	counter     uint64
	donecounter uint64
)

// Connection wraps in one struct the required parameters
// to connect to aws, region and AWS
type Connection struct {
	Region aws.Region
	Auth   aws.Auth
}

//Connect starts the connection to aws, and returns
// any error, default strategy is to connect 5 times, wait 4 seconds and with delay of 200 ms
// make sure to source the credential file on the server
// Region is harcoded to EUWest
// Auth is read from environment.
func (c *Connection) Connect() error {
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}
	c.Auth = auth
	c.Region = aws.EUWest
	return nil
}

// Getbucket connects to bucket name, with connection c, and returns
// a bucket type and any connection error
func Getbucket(name string, c *Connection) (*s3.Bucket, error) {
	s3Connector := s3.New(c.Auth, c.Region)
	bucket := s3Connector.Bucket(name)
	_, err := bucket.List("", "", "", 1)
	if err != nil {
		// probably wrong name of bucket
		log.Errorf("connection to s3", name, err)
	} else {
		fmt.Printf("Connected to s3 bucket: %v\n", name)
	}
	return bucket, err
}

// Put puts a serialize.Msgpack object into bucket b.
// by default uploads checksum of the files
func Put(b *s3.Bucket, data []byte, path string, checksum string) error {
	// data is a log of msgpack file kind
	// aws api wants []byte
	options := s3.Options{}
	options.ContentMD5 = checksum
	err := b.Put(path, data, "", s3.Private, options)
	if err != nil {
		return err
	}
	return nil
}

// Get gets the s3 object at the specified path and saves to disk
// with its filename and a root prependeded
func Get(b *s3.Bucket, path string, root string) error {
	var conn Connection
	err := conn.Connect()
	if err != nil {
		return err
	}
	data, err := b.Get(path)
	if err != nil {
		panic(fmt.Sprintf("can't download %v \n ", err))
	}
	localpath := fmt.Sprintf("%v/%v", root, path)
	stripped := strings.Split(localpath, "/")
	pathonly := strings.Join(stripped[:len(stripped)-1], "/")
	err = os.MkdirAll(pathonly, 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(localpath, data, 0666)
	if err != nil {
		return err
	}
	fmt.Printf("Downloaded: %v \n", localpath)
	return nil
}

// KeyBucket holds pointers to a bucket one wants to operate with
// and the key of interest inside that bucket
type KeyBucket struct {
	B   *s3.Bucket
	Key s3.Key
}

// Action is the the action on the key/bucket speficied in KeyBucket.
type Action func(KeyBucket)

//ApplyToMultiList applies the action a to all the keys that match prefix
//To select ALL contents of the bucket use prefix, delim = ""
func ApplyToMultiList(b *s3.Bucket, prefix, delim string, a Action) {
	//prints stats
	go func() {
		for {
			select {
			case <-time.After(3 * time.Second):
				r := atomic.LoadUint64(&donecounter)
				fmt.Printf("read %v keys\n", r)
				r = atomic.LoadUint64(&counter)
				fmt.Printf("processed %v keys\n", r)
			}
		}
	}()
	resp, err := b.List(prefix, delim, "", 1000)
	if err != nil {
		panic(err)
	}
	if len(resp.Contents) < 1 {
		log.Infof("got no Contents")
		return
	}
	lastSeen := resp.Contents[len(resp.Contents)-1]
	for _, obj := range resp.Contents {
		atomic.AddUint64(&counter, 1)
		go a(KeyBucket{b, obj})
	}
	for {
		if resp.IsTruncated {
			resp, err = b.List(prefix, delim, lastSeen.Key, 1000)
			if err != nil {
				panic(err)
			}
			lastSeen = resp.Contents[len(resp.Contents)-1]
			fmt.Printf("------ \n %v \n-----", lastSeen.Key)
			// TODO allow setting a max number of workers
			for _, obj := range resp.Contents {
				atomic.AddUint64(&counter, 1)
				go func() {
					wg.Add(1)
					a(KeyBucket{b, obj})
					wg.Done()
				}()
			}
		} else {
			break
		}
	}
	wg.Wait()
}

//Del deletes the key in the bucket
func Del(kb KeyBucket) {
	err := kb.B.Del(kb.Key.Key)
	if err != nil {
		log.Warnf("%v\n", err)
	}
	atomic.AddUint64(&donecounter, 1)
}
