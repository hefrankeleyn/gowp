package main

import (
	"fmt"
)

type GetIntChan func() <-chan int

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	intChan2 := GetIntChan(getIntChan)
	for elem := range intChan2() {
		fmt.Println(elem)
	}
}
