package action

import (
	mio "github.com/shopgun/matilde/io"
	"sync"
	"sync/atomic"
	"testing"
)

const basepath = "./foo/bar/2015/1/1/"

// MockReceive reads channel  and prints shizzzle
func testMockReceive(c chan NamedScanner, numFiles *uint64) {
	for _ = range c {
		atomic.AddUint64(numFiles, 1)
	}
}

// MockReceive reads channel  and prints shizzzle
func testMockReceiveLines(c chan mio.Input, numLines *uint64) {
	for _ = range c {
		atomic.AddUint64(numLines, 1)
	}
}

func TestFile(t *testing.T) {
	files := make(chan NamedScanner)
	done := make(chan struct{})
	paths := mio.WalkFiles(".", ".gz")
	var numFiles uint64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		SendFile(done, paths, files)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(files)
		close(done)
	}()
	testMockReceive(files, &numFiles)
	filestest := atomic.LoadUint64(&numFiles)
	if filestest != 2 {
		t.Errorf("expected 2 file got %v", filestest)
	}
}

func TestFileClosed(t *testing.T) {
	files := make(chan NamedScanner)
	done := make(chan struct{})
	close(done)
	paths := mio.WalkFiles(".", ".gz")
	var numFiles uint64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		SendFile(done, paths, files)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(files)
	}()
	testMockReceive(files, &numFiles)
	filestest := atomic.LoadUint64(&numFiles)
	// NOTE WTF this test fails at random!!!!
	if filestest > 3 {
		t.Errorf("expected   file at most 1 got %v", filestest)
	}
}

func TestLines(t *testing.T) {
	lines := make(chan mio.Input)
	files := make(chan NamedScanner)
	done := make(chan struct{})
	paths := mio.WalkFiles(".", ".gz")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		SendFile(done, paths, files)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(files)
	}()

	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		SendLine(done, files, lines)
		wg1.Done()
	}()
	go func() {
		wg1.Wait()
		close(lines)
		close(done)
	}()
	var numLines uint64
	testMockReceiveLines(lines, &numLines)
	linetest := atomic.LoadUint64(&numLines)
	if linetest != 6 {
		t.Errorf("expected 6 files got %v", linetest)
	}
}

func TestFileDoesnotExist(t *testing.T) {
	files := make(chan NamedScanner)
	done := make(chan struct{})
	paths := make(chan string, 1)
	paths <- "notexists"
	close(paths)
	var numFiles uint64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		SendFile(done, paths, files)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(files)
		close(done)
	}()
	testMockReceive(files, &numFiles)
	filestest := atomic.LoadUint64(&numFiles)
	if filestest != 0 {
		t.Errorf("expected 0 file got %v", filestest)
	}
}
