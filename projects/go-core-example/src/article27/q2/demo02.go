package main

import (
	"fmt"
	"strings"
)

func main() {
	var str01 string
	str01 = "Go，你好"
	// len(string) : 11
	fmt.Printf("len(string) : %v \n", len(str01))
	// len([]char): 5
	fmt.Printf("len([]char): %v\n", len([]rune(str01)))
	// len([]byte)： 11
	fmt.Printf("len([]byte)： %v\n", len([]byte(str01)))
	// 在string类型的值上进行切片操作，相当于对底层的字节数组做切片
	fmt.Println(str01[0:3]) // Go�
	fmt.Println(str01[0:5]) // Go，

	var sb01 strings.Builder
	sb01.WriteString("Go")
	fmt.Printf("Builder 的长度：%v\n", sb01.Len())
	sb01.Grow(10)
	fmt.Printf("Builder 扩容后的长度：%v\n", sb01.Len())

}
