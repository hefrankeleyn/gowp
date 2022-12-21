package main

import "fmt"

func modifyArraySlice(a [3][]string) [3][]string {
	s01 := a[0]
	if len(s01) > 0 {
		s01[len(s01)-1] = "x"
	}
	return a
}

func main() {
	a := [3][]string{
		[]string{"a", "b", "c"},
		[]string{"1", "2", "3"},
		[]string{"a1", "a2", "a3"},
	}
	fmt.Printf("开始值：%v\n", a)
	b := modifyArraySlice(a)
	fmt.Printf("修改后：%v\n", b)
	fmt.Printf("原始值：%v\n", a)
}
