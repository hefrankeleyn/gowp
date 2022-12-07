package main

import (
	"fmt"
)

func main() {
	// 示例1
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // 通道已满，这里会造成则色
	fmt.Println(<-ch1)

	// 示例2
	ch2 := make(chan int, 1)
	// elem, ok := <-ch2 // 通道已空，这里会造成阻塞
	// fmt.Printf("%d, %t\n", elem, ok)
	ch2 <- 1

	// 示例3
	var ch3 chan int
	// ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	// fmt.Println(<-ch3) // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3
}
