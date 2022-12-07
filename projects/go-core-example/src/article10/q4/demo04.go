package main

import (
	"fmt"
)

type Book struct {
	name  string
	price int
}

type Student struct {
	name string
	age  int
	desc []string
	oneB Book
}

func main() {
	ch01 := make(chan Student, 1)
	oneB := Book{name: "b01", price: 2}
	d01 := []string{"a", "b"}
	s01 := Student{name: "aa", age: 12, oneB: oneB, desc: d01}
	ch01 <- s01
	// 修改前的值：{name:aa age:12 desc:[a b] oneB:{name:b01 price:2}}
	fmt.Printf("修改前的值：%+v\n", s01)
	s01.name = "bb"
	s01.oneB.name = "b02"
	d01[1] = "c"
	// 修改后的值：{name:bb age:12 desc:[a c] oneB:{name:b02 price:2}}
	fmt.Printf("修改后的值：%+v\n", s01)
	// 从通道里出来的值：{name:aa age:12 desc:[a c] oneB:{name:b01 price:2}}
	fmt.Printf("从通道里出来的值：%+v\n", <-ch01)
}
