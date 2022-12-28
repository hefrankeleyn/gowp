package main

import "fmt"

func main() {
	// number2 := [...]int{1, 2, 3, 4, 5}
	number2 := []int{1, 2, 3, 4, 5}
	maxIndex := len(number2) - 1
	for i, e := range number2 {
		if i == maxIndex {
			number2[0] += e
		} else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
}
