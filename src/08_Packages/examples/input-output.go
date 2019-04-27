package main

import (
	"bytes"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("test"))
}