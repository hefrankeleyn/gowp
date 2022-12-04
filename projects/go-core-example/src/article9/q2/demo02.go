package main

import (
	"fmt"
)

/*
引起恐慌的示例
*/
func main() {
	var badMap2 = map[interface{}]int{
		"1":      1,
		3:        3,
		[]int{2}: 2, //  panic: runtime error: hash of unhashable type []int
	}
	fmt.Println(badMap2)
}
