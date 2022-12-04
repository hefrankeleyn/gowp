package main

import (
	"fmt"
)

/*
输出：
这个元素key "one"：1
*/
func main() {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "one"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("这个元素key %q：%d\n", k, v)
	} else {
		fmt.Println("没有发现！")
	}
}
