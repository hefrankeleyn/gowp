package main

import "fmt"

func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[1], value1[2]:
		fmt.Println("1 or 2")
	default:
		fmt.Println("other")
	}
}
