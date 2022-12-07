package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 准备好几个通道
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送数据
	rand.Seed(time.Now().Unix())
	index := rand.Intn(3)
	fmt.Printf("索引为:%d\n", index)
	intChannels[index] <- index
	// 哪个通道有可取元素，哪个队赢的分支就会执行
	select {
	case elm := <-intChannels[0]:
		fmt.Printf("第一个通道执行：%d\n", elm)
	case elm := <-intChannels[1]:
		fmt.Printf("第二个通道执行：%d\n", elm)
	case elm := <-intChannels[2]:
		fmt.Printf("第三个通道执行：%d\n", elm)
	default:
		fmt.Println("没有选中任何通道！")
	}
}
