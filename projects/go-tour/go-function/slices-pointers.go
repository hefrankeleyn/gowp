package main

import "fmt"

func main() {
	names := [4]string{
		"aa",
		"bb",
		"cc",
		"dd",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[1] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
