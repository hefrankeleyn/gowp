package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data"
	fmt.Printf("Write contents %q\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	// 7 bytes were read. (call Read)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	// The length of buffer: 31
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	// The capacity of buffer: 64
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
}
