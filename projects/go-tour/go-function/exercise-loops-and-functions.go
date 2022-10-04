package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var t float64
	for {
		// t = z
		// z -= (z*z - x) / (2 * z)
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-6 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(4))
}
