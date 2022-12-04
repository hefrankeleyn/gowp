package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个长度为5的环
	r := ring.New(5)
	// 获取环的长度
	n := r.Len()
	// 对环进行初始化
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// 倒序遍历环，并打印值
	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}
}
