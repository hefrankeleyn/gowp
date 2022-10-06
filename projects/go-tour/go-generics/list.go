package main

// List 代表一个单链表，可以持有任何数据
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) add(t T) {
	ol := l
	oneList := List[T]{ol, t}
	l = &oneList
}

func (l *List[T]) pop() T {
	var t T
	if l != nil {
		t = l.val
	}
	return t
}

func main() {
	// l := List.New[int]{"next": nil, "var": 23}
	// l.add(2)

	// fmt.Println()
}
