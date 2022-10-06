package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("两倍的%v是%v\n", v, v*2)
	case string:
		fmt.Printf("%q有%v个字节\n", v, len(v))
	default:
		fmt.Printf("我不知道的类型：%T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
