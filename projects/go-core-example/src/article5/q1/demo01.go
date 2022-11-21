package main

import (
	"fmt"
)

var name = "package"

func main() {
	name := "function"
	{
		name := "inner"
		fmt.Println(name)
	}
	fmt.Println(name)
}
