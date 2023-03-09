package main

import (
	"bufio"
	"strings"
)

func main() {
	reader01 := strings.NewReader("abcd")
	br01 := bufio.NewReader(reader01)
	p := make([]byte, 10)
	br01.Read(p)
	delim := byte(',')
	br01.ReadSlice(delim)
	br01.ReadBytes(delim)

}
