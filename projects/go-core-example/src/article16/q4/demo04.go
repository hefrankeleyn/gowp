package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	num := uint32(10)
	var count uint32 = 0
	trigger := func(i uint32, fn func()) {
		for {
			if atomic.LoadUint32(&count) == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			// 这里加Sleep语句是很有必要的
			time.Sleep(time.Microsecond)
		}
	}
	for i := uint32(0); i < num; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	fmt.Printf("goroutine数量：%v\n", runtime.NumGoroutine())
	trigger(num, func() {})
}
