package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个长度为5 的环
	r := ring.New(5)
	// 获取ring的长度
	n := r.Len()
	// 对r 的值进行初始化
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// 将指针向前移动三步
	// r = r.Move(3)
	r = r.Move(-3)

	// 迭代ring
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
