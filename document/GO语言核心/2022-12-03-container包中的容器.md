# container 包中的容器

[toc]

## 一、container/list 双向链表

### 1.1 参考：[container/list](https://pkg.go.dev/container/list)

在container/list代码包中，有两个公开的程序实体——List和Element。List实现了一个双向链表，而Element则代表链表中元素的结构。

Element中包含的值是一个[`interface{}`类型](https://pkg.go.dev/builtin#any)。

### 1.2 List和Element的使用示例

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	oneList := list.New()
	// fmt.Printf("oneList： %T, %v\n", oneList, oneList.Len())
	e1 := oneList.PushBack(4)
	// fmt.Printf("e1： %T, %v\n", e1, e1.Value)
	e2 := oneList.PushFront(1)
	// fmt.Printf("oneList: %v, e2： %T, %v\n", oneList, e2, e2.Value)
	oneList.InsertAfter(5, e1)
	oneList.InsertBefore(8, e2)
	for e := oneList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
```

## 二、container/ring 循环链表

- `func New(n int) *Ring` 创建一个长度为n的循环链表
- `func (r *Ring) Len() int`  获取循环链表的长度
- `func (r *Ring) Do(f func(any))`  按照向前的顺序，在每一个ring的元素上调用函数f。
- `func (r *Ring) Next() *Ring` 返回下一个环的元素，r必须不能为空。
- `func (r *Ring) Prev() *Ring` 返回下一个环的元素。

```go
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
```

- `func (r *Ring) Link(s *Ring) *Ring`  链接两个环

  - r 不能为空。

  - 如果r和s指向相同的环，链接的环，将移除r和s之间的元素。

```go
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
	s := ring.New(2)
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
```

- `func (*Ring) Move(n int) *Ring`

  在环中向前或向后移动`n%r.Len()`元素，并返回该环元素，必须不能为空。
  
  ```go
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
  
  ```

- `func (*Ring) Unlink(n int) *Ring`

  从环中，移除n%r.Len() 个元素，从r.Next() 开始。如果n%r.len()==0，r保持不变

  ```go
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
  ```

## 三、container/heap

### 3.1 参考[container/heap](https://pkg.go.dev/container/heap)

heap 包为任何实现了heap.interface 接口的类型提供堆操作。一个堆就是一个树，这个树的特点是每个节点在它子树中都是最小值节点。

在树中，最小值的元素是root，它的索引为0。

### 3.2 示例一：IntHeap

这个示例演示了一个整数类型的堆，通过实现heap.interface。

```go
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
```

### 3.3 示例二：PriorityQueue 

这个示例使用一些项，创建一个优先级队列，添加并维护一个项，然后按照顺序从优先级队列中移除。

```go
package main

// 这个示例演示通过使用heap 接口，构建一个优先级队列
import (
	"container/heap"
	"fmt"
)

// 一个我们在优先级队列中管理的项
type Item struct {
	value    string // 项目的值
	priority int    // 项目在队列中的优先级
	// 索引是需要被更新，它是被heap.interface 的方法维护
	index int // 项目在堆中的索引
}

// 一个优先级队列，实现了heap.interface 并持有Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// 我们想让Pop给我们使用最大，而不是最小，优先级，因此使用更大比较
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := (*pq).Len()
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := (*pq).Len()
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // 为了安全
	*pq = old[0 : n-1]
	return item
}

// update 修改优先级 和 在队列中 Item 的值
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// 这个示例创建一个带有若干项的PriorityQueue，添加并操作一个item，
// 然后按照顺序从优先级队列中移除元素
func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	// 创建一个优先级队列，并把项目放到它里面，然后建立优先级队列的不变量
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	// 插入一个新元素并修改它的优先级
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	// 取出物品，他们到达，按照优先级顺序
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d : %s\n", item.priority, item.value)
		// fmt.Printf("%2.d : %s\n", item.priority, item.value)
	}
}
```



