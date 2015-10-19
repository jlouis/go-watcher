package serialize

import (
	"bufio"
	"bytes"
	"io"

	"github.com/giulioungaretti/structs"
	"github.com/ugorji/go/codec"
)

// Msgpack wraps around the codec/msgpack struct
// to provide a init step
type Msgpack struct {
	Init bool
	E    *codec.Encoder
}

// Init is required to specifiy which reader the msgpacks should be written to
func Init(w io.Writer) Msgpack {
	handle := new(codec.MsgpackHandle)
	encoder := codec.NewEncoder(w, handle)
	m := Msgpack{}
	m.E = encoder
	m.Init = true
	return m
}

// Serialize serialize an interface (which must be struct) into msgpack
// using the encoder m. Field names are lowercased.
func (m *Msgpack) Serialize(in interface{}) (err error) {
	s := structs.New(in)
	lowerCased := s.LowerCaseMap() // Get a map[string]interface{}
	err = m.E.Encode(lowerCased)
	return
}

// Encode encodes the in interface into a msgpack streamed to the file  or
// into a bytes buffer configured inside the encoder struct
func Encode(in interface{}) (msgpack []byte, err error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	h := new(codec.MsgpackHandle)
	Encoder := codec.NewEncoder(w, h)
	s := structs.New(in)
	lowerCased := s.LowerCaseMap() // Get a map[string]interface{}
	err = Encoder.Encode(lowerCased)
	w.Flush()
	msgpack = b.Bytes()
	return
}
