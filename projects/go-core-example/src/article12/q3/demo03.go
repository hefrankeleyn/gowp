package main

import (
	"errors"
	"fmt"
)

type calculateFunc func(x, y int) (int, error)

type operate func(x, y int) int

func genCalculate(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("无效的操作")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y := 58, 59
	op := func(x, y int) int {
		return x + y
	}
	add := genCalculate(op)

	res, err := add(x, y)
	fmt.Printf("结果为：%d （error：%v）\n", res, err)
}
