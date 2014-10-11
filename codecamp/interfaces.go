package main

import (
	"bytes"
	"io"
	"os"
)

// START OMIT
type Shouty struct {
	w io.Writer
}

func (s Shouty) Write(data []byte) (n int, err error) {
	return s.w.Write(bytes.ToUpper(data))
}

func main() {
	var w io.Writer = Shouty{os.Stdout}
	w.Write([]byte("Hello gophers!"))
}

// END OMIT
