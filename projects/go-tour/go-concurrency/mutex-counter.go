package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 在并发中可以安全使用
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// 对给定key进行递增
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// 锁上，以便某一个时刻只有一个goroutine能读取 map c.v
	c.v[key]++
	c.mu.Unlock()
}

// 获取给定key的统计值
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// 锁上，以便某一个时刻只有一个goroutine能读取 map c.v
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}
