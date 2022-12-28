package main

import "fmt"

func main() {
	num := 10
	var ch = make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < num; i++ {
		<-ch
	}
}
