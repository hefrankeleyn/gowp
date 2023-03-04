package main

import (
	"bytes"
	"fmt"
)

func main() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	// The capacity of new buffer with contents "ab": 8
	// 容量为何为8，看 runtime/string.go#stringtoslicebyte()
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", contents, buffer1.Cap())
	unreadBytes := buffer1.Bytes()
	// The unread bytes of the buffer: [97 98]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	buffer1.WriteString("cdefg")
	// The capacity of new buffer with contents "ab": 8
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", contents, buffer1.Cap())
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	// 基于前面的内容获取到结果值
	// The unread bytes of the buffer: [97 98 99 100 101 102 103 0]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	// 操纵buffer
	unreadBytes[len(unreadBytes)-2] = byte('X')
	// The unread bytes of the buffer: [97 98 99 100 101 102 88 0]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
}
