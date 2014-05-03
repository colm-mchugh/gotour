package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt *rot13Reader) Read(p []byte) (n int, err error) {
	var numread int
	var errcode error

	numread, errcode = rt.r.Read(p)
	for i := 0; i < numread; i++ {
		switch {
        case ((p[i] >= 'a') && (p[i] <= 'm')):
            fallthrough
        case ((p[i] >= 'A') && (p[i] <= 'M')):
            p[i] += 13
        case ((p[i] >= 'n') && (p[i] <= 'z')):
            fallthrough
        case ((p[i] >= 'N') && (p[i] <= 'Z')):
            p[i] -= 13
		}
	}
	return numread, errcode
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
