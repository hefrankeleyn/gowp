package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{4, 3}

	a = f  // MyFloat 实现了 Abser
	a = &v // 一个 *Vertex 实现了 Abser

	// 在这一行，v是一个Vertex（不是一个 *Verterx），并且没有实现 Abser
	a = v

	fmt.Println(a.Abs())
}
