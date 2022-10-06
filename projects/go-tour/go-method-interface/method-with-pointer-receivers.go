package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("在调用Sacle之前：%v，Abs：%v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("在调用Scale之后：%+v, Abs：%v\n", v, v.Abs())
}
