package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// coordinateWithWaitGroup()
	coordinateWithContext()
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			// 如果所有的addNum函数都执行完毕，那么就立即分发子任务的goroutine
			// 这里分发子任务的goroutine，就是执行 coordinateWithContext 函数的goroutine.
			if atomic.LoadInt32(&num) == int32(total) {
				// <-cxt.Done() 针对该函数返回的通道进行接收操作。
				// cancelFunc() 函数被调用，针对该通道的接收会马上结束。
				// 所以，这样做就可以实现“等待所有的addNum函数都执行完毕”的功能
				cancelFunc()
			}
		})
	}
	<-cxt.Done()
	fmt.Println("end.")
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i += stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
}

func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		// time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			fmt.Printf("The CAS option failed. [%d-%d]\n", id, i)
		}
	}
}
