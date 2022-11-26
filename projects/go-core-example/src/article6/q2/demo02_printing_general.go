package main

import (
	"fmt"
)

type School struct {
	name string
	s_id int
}

func main() {
	example := "hello, world"
	example1 := School{name: "s01", s_id: 12}
	// %v 示例1：hello, world， 示例二:{s01 12}
	fmt.Printf("%%v 示例1：%v， 示例二:%v\n", example, example1)
	// %+v 示例1：hello, world， 示例二:{name:s01 s_id:12}
	fmt.Printf("%%+v 示例1：%+v， 示例二:%+v\n", example, example1)
	// %T, 示例1：string, 示例二：main.School
	fmt.Printf("%%T, 示例1：%T, 示例二：%T\n", example, example1)
	// %%
	fmt.Printf("百分号的字面量，不会消耗值%%v\n")
	example2 := true
	// %t 一个为ture或false的单词
	// ture
	fmt.Printf("%t\n", example2)
}
