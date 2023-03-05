package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	src := strings.NewReader(
		"CopyN copies n bytes (or unil an error) from src to dst." +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		// Written(58): "CopyN copies n bytes (or unil an error) from src to dst.It"
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
	// teeReader1 := io.TeeReader(src, dst)
	// p := make([]byte, 7)
	// teeReader1.Read(p)
	// io.MultiReader()
	// io.Pipe()
}
