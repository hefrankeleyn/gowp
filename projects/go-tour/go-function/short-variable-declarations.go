package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java, i1 := true, false, "no!", 23.1
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("%T : %v\n", i1, i1)
	var i2 = 21
	fmt.Printf("%T : %v\n", i2, i2)
}
