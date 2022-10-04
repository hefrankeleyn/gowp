package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向i
	fmt.Println(*p) // 通过指针读取i
	*p = 21         // 通过指针设置i
	fmt.Println(i)  // 查看i的新值

	p = &j         // 指向j
	*p = *p / 37   // 通过指针除j
	fmt.Println(j) // 查看j的新值
}
