package action

import (
	"bufio"
	"compress/gzip"
	"expvar"
	mio "github.com/shopgun/matilde/io"
	"os"

	log "github.com/Sirupsen/logrus"
)

// NamedScanner contains a pointer to the buffer we want
// to scan and the path from where we readed it from.
type NamedScanner struct {
	Path    string
	Scanner *bufio.Scanner
}

//holds total number of events processed
var (
	batchCounts = expvar.NewInt("Total events batch")
	fileCounts  = expvar.NewInt("Processed files")
)

// SendFile reads gz filenames and sends a scanner through the channel
// so that a consumer can scan through its lines
func SendFile(done <-chan struct{}, paths <-chan string, files chan NamedScanner) {
	for path := range paths {
		//load dirty gizp file
		file, err := os.Open(path)
		if err != nil {
			log.Errorf("io error %v for %v", err, path)
			//go to next file
			continue
		}
		raw, err := gzip.NewReader(file)
		if err != nil {
			log.Errorf("io error %v for %v", err, path)
			//go to next file
			continue
		}
		scanner := bufio.NewScanner(raw)
		scanner.Split(bufio.ScanLines)
		select {
		case files <- NamedScanner{
			Path:    path,
			Scanner: scanner,
		}:
		case <-done:
			return
		}
		fileCounts.Add(1)
		log.Debugf("sent:%v\n", path)
	}
}

// SendLine reads a file and sends lines thorugh channel
func SendLine(done <-chan struct{}, files chan NamedScanner, lines chan mio.Input) {
	for namedscanner := range files {
		scanner := namedscanner.Scanner
		path := namedscanner.Path
		ty := mio.GetType(path)
		scanner.Split(bufio.ScanLines)
		log.Debugf("Start sending line from  %v", path)
		for scanner.Scan() {
			b := scanner.Bytes()
			line := make([]byte, len(b))
			copy(line, b)
			event := mio.Input{
				Content: line,
				Type:    ty,
			}
			select {
			case lines <- event:
			case <-done:
				return
			}
			batchCounts.Add(1)
		}
		log.Debugf("Finished sending lines from %v", path)
	}
}
