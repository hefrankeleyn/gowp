package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// 遍历树t，发送所有的值到 ch
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 判断 t1 和 t2 是否包含相同的值
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		x, y := <-c1, <-c2
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	// t := tree.New(2)
	// ch := make(chan int, 10)
	// go Walk(t, ch)

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v ", <-ch)
	// }
	// fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
