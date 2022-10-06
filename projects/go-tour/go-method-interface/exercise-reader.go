package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	// b = append(b, 'A')
	return 1, nil
}

func main() {
	b := make([]byte, 0)
	b2 := byte(int('A') + 13)
	fmt.Printf("%q\n", b2)
	b = append(b, 'A')
	fmt.Printf("%q\n", b)
	reader.Validate(MyReader{})

}
