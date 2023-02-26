package main

import (
	"fmt"
	"strings"
)

func main() {
	var reader1 strings.Reader
	reader1 = *strings.NewReader("Go， 你好")
	ch, size, err := reader1.ReadRune()
	fmt.Printf("%v, %v, %v\n", string(ch), size, err)
	readingIndex := reader1.Size() - int64(reader1.Len())
	fmt.Printf("以读计数：%v\n", readingIndex)
}
