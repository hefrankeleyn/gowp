package main

import (
	"fmt"
)

var channels = [3]chan int{
	nil,
	make(chan int, 1),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("第一种情况被选择!")
	case getChan(1) <- getNumber(1):
		fmt.Println("第二种情况被选择!")
	case getChan(2) <- getNumber(2):
		fmt.Println("第三种情况被选择!")
	default:
		fmt.Println("没有情况被选中！")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
