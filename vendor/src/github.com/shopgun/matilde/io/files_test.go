package io

import (
	"os"
	"strconv"
	"sync/atomic"
	"testing"

	log "github.com/Sirupsen/logrus"
)

const path string = "./data/"

const badpath string = "/Users/giulio/foo/bar"

func TestGood(t *testing.T) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		// makse sure we have data
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	filesno := 200
	for i := 0; i < filesno; i++ {
		name := path + strconv.Itoa(i) + "test.gz"
		if _, err := os.Stat(name); !os.IsExist(err) {
			fs, err := os.Create(name)
			if err != nil {
				log.Fatal(err)
			}
			fs.Close()
		}
	}
	channel := WalkFiles(path, ".gz")
	//make sure we get data
	var files uint64
	for _ = range channel {
		atomic.AddUint64(&files, 1)
	}
	total := atomic.LoadUint64(&files)
	if int(total) != filesno {
		t.Error("Got % files, wanted % files", total, filesno)
	}
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}

func TestScopeAction(t *testing.T) {
	path := "offer/click/2015/1/1/dummy.gz"
	scope := GetScope(path)
	action := GetAction(path)
	if scope+action != "offerclick" {
		t.Errorf("Parsing path incorrectly")
	}
}
