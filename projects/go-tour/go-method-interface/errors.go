package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("当前时间 %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "它不能工作"}
}

func main() {
	fmt.Println(math.Abs(3 - 1))
	fmt.Println(math.Abs(2-1) < 1e-6)
	fmt.Println(1e-6 < 0.000001)
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
