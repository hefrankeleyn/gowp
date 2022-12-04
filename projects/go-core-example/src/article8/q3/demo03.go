package main

import (
	"container/ring"
	"fmt"
)

/*
输出：
0
0
1
1
*/
func main() {
	//  创建两个循环链表，r 和 s， 长度都为2
	r := ring.New(2)
	// s := ring.New(2)
	s := r
	//  获取两个循环链表的长度
	lr := r.Len()
	ls := s.Len()

	// 将r 的值初始化为 0
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// 将s 的值初始化为 1
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}
	// 链接循环链表 r 和 s
	rs := r.Link(s)

	// 迭代组合的环， 并打印它们的值
	rs.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
