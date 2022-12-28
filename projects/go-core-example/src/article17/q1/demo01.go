package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 5, 6, 7}
	for i := range number1 {
		if i == 3 {
			number1[i] |= i
		}
	}
	fmt.Println(number1)
}
