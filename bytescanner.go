package bytescanner

import (
	"io"
)

// ByteScanner represents an instance of a Byte Scanner bound to one io.Reader.
type ByteScanner struct {
	reader io.Reader
	// [...]
}

// NewByteScanner returns a byte scanner that takes a io.Reader as a source, and
// that provides an Eat() and Peek() methods.
// The Eat() method returns the first byte of the remaining bytes (and removes
// it from the remaining bytes to scan).
// The Peek(n) return the n'th bytes of the remaining bytes.  Calling Peek(0)
// will return the byte that Eat() would return (without removing it from the
// remaining bytes to scan), while Peek(1) will return the following byte, etc.
// The Peek(n) method handles stepping over the io.Reader buffer boundaries and,
// like Eat(), handles io.EOF properly.
func NewByteScanner(reader io.Reader) *ByteScanner {
	return &ByteScanner{
		reader: reader,
		// [...]
	}
}

// Eat returns the first byte of the remaining bytes from the io.Reader stream.
// Calling it again will return the next byte, and so on until io.EOF.
// Calling Eat on an empty io.Reader will return an io.EOF error.
func (bs *ByteScanner) Eat() (byte, error) {
	// [...]
	return 0, io.EOF
}

// Peek returns one of the remaining bytes without removing it from the
// stream (of remaining bytes).
// Calling Peek(n) return the n'th bytes from the remaining byte stream.
// Calling Peek(0) will return the byte that Eat() would return (without
// removing it from the remaining bytes to scan), while Peek(1) will return the
// following byte, etc. The Peek(n) method handles stepping over the io.Reader
// buffer boundaries and, like Eat(), handles io.EOF properly.
func (bs *ByteScanner) Peek(offset int) (byte, error) {
	// [...]
	return 0, io.EOF
}
