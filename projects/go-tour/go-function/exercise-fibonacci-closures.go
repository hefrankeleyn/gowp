package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 0
	return func() int {
		if first == 0 {
			first = 1
			second = 1
			return 0
		} else {
			current := first
			firstc := second
			second = first + second
			first = firstc
			return current
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
