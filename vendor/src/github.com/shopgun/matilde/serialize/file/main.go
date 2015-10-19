package serialize

import (
	"fmt"
	"os"
	"strings"

	xz "github.com/giulioungaretti/go-xz"
	"github.com/giulioungaretti/structs"
	"github.com/ugorji/go/codec"
)

// InitMsgPack initializes the msgpack encoder
// accepts path as in the path where the log of msgpacks
// are stored. Returns error when a)can't create directory structure or
// b) the file itself.
func (m *Msgpack) InitMsgPack(path string, base string) error {
	//
	h := new(codec.MsgpackHandle)
	// path ninja fu
	msgpackPath := strings.Replace(path, ".json", "", 1)
	// same path with correct extension
	msgpackPath = strings.Replace(msgpackPath, ".gz", "", 1)
	// extract relative path
	pathSplit := strings.Split(msgpackPath, "/")
	relativepathSplit := pathSplit[len(pathSplit)-6 : len(pathSplit)]
	// relative path of the msgpack
	relativepath := strings.Join(relativepathSplit, "/")
	// replace / with _ so it's a flat directory structure
	relativepath = strings.Replace(relativepath, "/", "_", -1)
	// save to temp dir
	base = fmt.Sprintf("%v/tmp", base)
	// now join with new base path
	newmsgpackPath := strings.Join([]string{base, relativepath}, "/")
	var err error
	var msgpackFile *os.File
	switch base {
	default:
		// first create subdirs
		err := os.MkdirAll(base, 0777)
		if err != nil {
			return err
		}
		m.Path = newmsgpackPath
	case "debug":
		//save to root folder
		filename := pathSplit[len(pathSplit)-1]
		path := fmt.Sprintf("./%v", filename)
		m.Path = path
	case "":
		m.Path = msgpackPath
	}

	msgpackFile, err = os.OpenFile(m.Path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		// panic if local file can't be created
		return err
	}
	m.File = msgpackFile
	m.Encoder = codec.NewEncoder(m.File, h)
	return err
}

// LogMsgPack encodes the in interface into a msgpack streamed to the file  or
// into a bytes buffer configured inside the encoder struct
func (m *Msgpack) LogMsgPack(in interface{}) error {
	s := structs.New(in)
	lowerCased := s.LowerCaseMap() // Get a map[string]interface{}
	err := m.Encoder.Encode(lowerCased)
	return err
}

// Closelog  makes sure that the file underneath the encoder struct
// gets closed
func (m *Msgpack) Closelog() string {
	m.File.Close()
	outname := fmt.Sprintf("%v.msgpack", m.Path)
	outname = strings.Replace(outname, "/tmp", "", 1)
	os.Rename(m.Path, outname)
	return outname
}

// Msgpack contains the details about the path  of the log file for the msgpacks
// the file in memory, and the encoder in memory
// this is just a wrapper but makes very easy to tweak the encoder with  the initMsgPack
// method.
type Msgpack struct {
	File     *os.File
	Bytes    []byte
	Encoder  *codec.Encoder
	Path     string
	Checksum xz.Checksum
}
