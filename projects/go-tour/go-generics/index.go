package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 都是T类型，它们有用comparable约束，因为能在这里使用 ==
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index 能作用在int类型的切片上
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index 也能作用在string类型的切片行
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "Hello"))
}
