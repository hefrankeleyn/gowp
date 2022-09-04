package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed agained: %q, err: %v\n", doubleRev, doubleErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("输入是一个无效的UTF-8")
	}
	// fmt.Printf("input: %q\n", s)
	r := []rune(s)
	// fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)-1; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
