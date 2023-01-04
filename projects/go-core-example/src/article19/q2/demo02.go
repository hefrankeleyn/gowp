package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main")
	// 延迟func函数的执行，直到main结束
	defer func() {
		fmt.Println("Enter defer function")
		if p := recover(); p != nil {
			fmt.Printf("%v\n", p)
		}
		fmt.Println("Exit defer function")
	}()
	// 引发painc
	panic(errors.New("soming wrong"))
	fmt.Println("Exit function main")
}
