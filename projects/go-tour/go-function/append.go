package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append 使用于 nil 切片
	s = append(s, 0)
	printSlice(s)

	// 切片根据需要增长
	s = append(s, 1)
	printSlice(s)

	// 一次能够添加多个元素
	s = append(s, 2, 3, 4, 5)
	printSlice(s)

	s2 := append(s, 10)
	printSlice(s2)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), len(s), s)
}
