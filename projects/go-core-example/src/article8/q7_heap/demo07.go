package main

import (
	"container/heap"
	"fmt"
)

// 一个IntHeap是最小整数的小顶堆
type IntHeap []int

// 实现 sort.interface 接口方法
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现heap.interface 的 Push方法
func (h *IntHeap) Push(x any) {
	// Push 和 Pop 使用指针接收器，因为它们不仅修改了它的内容，还修改了切片的长度。
	*h = append(*h, x.(int))
}

// 实现heap.interface 的 POP方法
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 这个示例向IntHeap中插入若干个int值，检查最小值，并按照优先级顺序移除它们
/*
输出：
minimum：1
1 2 3 5
*/
func main() {
	// v01 := []int{1, 2, 3}
	// fmt.Println(v01)
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum：%d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
