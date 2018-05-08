package main

import "io"

// Counter takes a io.Reader an a pointer to a int counter and keeps track of the written bytes of R
type Counter struct {
	R io.Reader
	N *int
}

// Read satisfies the io.Reader interface and increments the counter
func (c Counter) Read(p []byte) (n int, err error) {
	n, err = c.R.Read(p)

	*c.N += n

	return n, err
}
