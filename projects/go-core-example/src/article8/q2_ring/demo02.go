package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个大小为5的循环链表
	r := ring.New(5)
	// 获取循环链表的长度
	n := r.Len()
	fmt.Printf("循环链表的长度：%v\n", n)
	// 对这个循环链表初始化整数类型的值
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// 迭代这个循环链表，打印它的值
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})
}
