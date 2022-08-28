package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// 为Int 值初始化一个map
	ints := map[string]int64{
		"first":  34,
		"second": 21,
	}
	// 为float值 初始化一个map
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("非泛型 Sums：%v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("泛型求和： %v 和 %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("泛型求和，推断类型参数： %v 和 %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("泛型求和，使用约束： %v 和 %v\n", SumNumbers(ints), SumNumbers(floats))
}

// SumInts 将m的值添加在一起
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats 将m的值添加在一起
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloat 获取m值的和。它既支持int64类型，也支持float64类型
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers 对m的值求和。它既支持integer，又支持floats 作为map的值
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
