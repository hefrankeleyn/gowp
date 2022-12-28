package main

import "fmt"

func main() {
	num := 3
	switch num {
	case 3:
		fmt.Println("3 case")
		fallthrough
	case 2:
		fmt.Println("2 case")
	case 1:
		fmt.Println("1 case")
	default:
		fmt.Println("no case")
	}
}
