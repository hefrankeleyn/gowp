package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 这个方法意味着类型T实现了接口I，
// 但是不需要明确的声明
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
