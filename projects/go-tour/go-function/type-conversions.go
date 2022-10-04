package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 42
	var f1 float64 = float64(i)
	var u1 uint = uint(f1)
	fmt.Println(i, f1, u1)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
