package main

import (
	"errors"
	"fmt"
)

type operate func(a, b int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("无效的操作")
	}
	return op(x, y), nil
}

func main() {
	op := func(a, b int) int {
		return a + b
	}

	x, y := 2, 1
	res, err := calculate(x, y, op)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

}
