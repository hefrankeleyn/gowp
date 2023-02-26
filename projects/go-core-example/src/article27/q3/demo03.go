package main

import "strings"

func main() {
	var sb01 strings.Builder
	sb01.WriteString("Go")
	sb02 := sb01
	// sb02.Grow(1) // 这里会引起恐慌panic
	_ = sb02

	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1)
	}
	f2(&sb01)

	sb01.Reset()
	sb03 := sb01
	sb03.Grow(1)
}
