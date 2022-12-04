package main

import (
	"container/ring"
	"fmt"
)

/*
输出：
0
4
*/
func main() {
	// 创建一个长度为5的环
	r := ring.New(5)
	// 获取ring的长度
	n := r.Len()
	// 初始化ring的值
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// 从r.Next() 开始，移除 3 个元素
	r.Unlink(3)
	// 迭代剩余的环，打印它的内容
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})
}
