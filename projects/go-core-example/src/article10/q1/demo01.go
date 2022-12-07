package main

import (
	"fmt"
)

/*
输出：
1
*/
func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	fmt.Printf("管道ch1 接收的第一个元素：%d\n", <-ch1)
}
