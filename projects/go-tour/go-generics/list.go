package main

import "fmt"

// List 代表一个单链表，可以持有任何数据
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(t T) {
	if l.next == nil {
		l.next = &List[T]{val: t}
	} else {
		l.next.Push(t)
	}
}

func (l *List[T]) GetALL() []T {
	res := make([]T, 0)
	l2 := l
	for l2 != nil {
		res = append(res, l2.val)
		l2 = l2.next
	}
	return res
}

func main() {
	l := List[int]{val: -1}
	oneL := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range oneL {
		l.Push(v)
	}
	fmt.Println(l.GetALL())
	// l := List.New[int]{"next": nil, "var": 23}
	// l.add(2)

	// fmt.Println()
}
