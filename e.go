package main

import (
	"io"
	"os"
)

type eReader struct{}

func (eReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'e'
	}
	return len(b), nil
}

func main() {
	io.Copy(os.Stdout, eReader{})
}
