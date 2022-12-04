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
	heap.Remove(&pq, 1)
	// 取出物品，他们到达，按照优先级顺序
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d : %s, index: %d\n", item.priority, item.value, item.index)
		// fmt.Printf("%2.d : %s\n", item.priority, item.value)
	}
}
