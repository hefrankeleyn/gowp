package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// 切片这个切片给一个0长度
	s = s[:0]
	printSlice(s)

	// 扩展它的长度
	s = s[:6]
	printSlice(s)

	// 删除它前两个值
	s = s[2:6]
	printSlice(s)
	// s = s[:6]
	// printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
