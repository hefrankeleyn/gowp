package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道
	time.AfterFunc(time.Second*5, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("这个计数器是关闭了！")
			break
		}
		fmt.Println("计数器还在运行！")
	}

}
