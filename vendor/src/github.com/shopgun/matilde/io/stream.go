package io

const (
	// Keybytes defines the byte size of the Key length
	Keybytes = 1
)

// Input defines the structure of input data. It is given by two fields,
// the Content which is the raw content and Type which signifies what
// kind of content we are working with
type Input struct {
	Content []byte
	Type    []byte
}
