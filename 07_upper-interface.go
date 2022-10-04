package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Upper struct {
	upr io.Writer
}

func (u *Upper) Write(b []byte) (n int, err error) {
	out := bytes.ToUpper(b)
	return u.upr.Write(out)
}

func main() {
	c := &Upper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
