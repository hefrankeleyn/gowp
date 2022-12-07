package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	// 发送方
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("发送方：发送元素：%v\n", i)
			ch1 <- i
		}
		fmt.Println("发送方：关闭通道")
		close(ch1)
	}()

	// 接收方
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("接收方：元素接收完毕！")
			break
		}
		fmt.Printf("接收方：接收到元素%d\n", elem)
	}

	fmt.Println("结束。")
}
