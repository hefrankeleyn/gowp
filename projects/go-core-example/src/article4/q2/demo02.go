package main

import (
	"fmt"
	"io"
	"os"
)

var name = "cc"

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
	num01, name := 01, "aa"
	fmt.Println(num01, name)
}
