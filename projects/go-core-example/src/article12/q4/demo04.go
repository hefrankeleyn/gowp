package main

import "fmt"

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("数组：%v\n", array1) // 数组：[a b c]
	array2 := modifyArray(array1)
	fmt.Printf("修改后：%v\n", array2) // 修改后：[a x c]
	fmt.Printf("原始值：%v\n", array1) //  原始值：[a b c]
}
