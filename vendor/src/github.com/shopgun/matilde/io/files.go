package io

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	// Delimiter that is used in the file system by default
	Delimiter = "/"
)

// WalkFiles starts a goroutine that send to  channel the paths  starting from root recursively, of files
// whose extension matches ext, NOTE hidden files included.
// Errors are send thorugh the errorc channel, when done is closed, then no more work is done.
// TODO don't forget about the error channel.
func WalkFiles(root string, ext string) <-chan string {
	pathc := make(chan string)
	go func() {
		// Close the paths/errc channel after Walk returns.
		defer close(pathc)
		// Visit fucntion for walk
		// returns path for file if file extension
		// matches ext. F.ex. ".gz", ".zip"
		var visit = func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ext && !f.IsDir() {
				select {
				case pathc <- path:
				}
			}
			// no error returned
			return nil
		}
		// send value through error channel
		filepath.Walk(root, visit)
	}()
	return pathc
}

// AFile is a simple auxillary struct that holds a path to a file and any error
// that may have occured when loading it
type AFile struct {
	Path string
	Err  error
}

// GetScope extracts the scope of the event from the path of the file
func GetScope(path string) string {
	components := strings.Split(path, Delimiter)
	maxl := len(components)
	scope := components[maxl-6]
	return scope
}

// GetAction extracts the scope of the event from the path of the file
func GetAction(path string) string {
	components := strings.Split(path, Delimiter)
	maxl := len(components)
	action := components[maxl-5]
	return action
}

// GetType returns the type of the event at path.
func GetType(path string) []byte {
	scope := GetScope(path)
	action := GetAction(path)
	eventType := fmt.Sprintf("%v.%s", scope, action)
	return []byte(eventType)
}

// LoadCsv loads a csv files with "fields" number of fields
// returns only strings version of the fields of the csv
func LoadCsv(path string, fields int) ([][]string, error) {
	csvfile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reader := csv.NewReader(csvfile)
	reader.Comma = '\t' // Use tab-separated values
	// For speed and flexibiliy just load the
	// fields we need
	reader.FieldsPerRecord = fields
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	csvfile.Close()
	return rawCSVdata, nil
}

// AppendLineToTxt appends a single line/string to a file
// concurrent writes safe
func AppendLineToTxt(path string, line string) error {
	var mutex = &sync.Mutex{}
	// only for appending
	mutex.Lock()
	// this is gizpedd
	fi, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	fiwriter := bufio.NewWriter(fi)
	if err != nil {
		return err
	}
	fiwriter.WriteString(line)
	fiwriter.Flush()
	fi.Close()
	mutex.Unlock()
	return err
}
