package main

import (
	"fmt"
)

func main() {
	s6 := make([]int, 0)
	fmt.Printf("切片s6的容量为：%d\n", cap(s6))
	for i := 1; i <= 5; i++ {
		// 扩容切片，或生成一个新的切片
		s6 = append(s6, i)
		// 一般情况下， 每次扩容，如果容量不够， 新切片的长度是原切片的 2倍
		fmt.Printf("s6(%d)： 长度：%d， 容量：%d\n", i, len(s6), cap(s6))
	}
	fmt.Println()
	// 示例二：
	// 如果原始切片的容量大于等于1024， 一旦需要扩容，会以原容量的1.5倍作为新容量的的基准，不断调整（不断乘以1.5），
	// 直到结果不小于原长度与要追加长度的元素数量之和
	s7 := make([]int, 1024)
	fmt.Printf("切片s7的容量为：%d\n", cap(s7))
	s7 = append(s7, make([]int, 200)...)
	fmt.Printf("s7长度：%d， 容量：%d\n", len(s7), cap(s7))
	s7 = append(s7, make([]int, 400)...)
	fmt.Printf("s7长度：%d， 容量：%d\n", len(s7), cap(s7))
	s7 = append(s7, make([]int, 800)...)
	fmt.Printf("s7长度：%d， 容量：%d\n", len(s7), cap(s7))
	s7 = s7[0 : cap(s7)-800]
	fmt.Printf("s7长度：%d， 容量：%d\n", len(s7), cap(s7))
}
