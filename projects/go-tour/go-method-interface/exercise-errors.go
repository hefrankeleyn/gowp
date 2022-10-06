package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintln("cannot Sqrt negative number:", e) // 会造成无限循环
	return fmt.Sprintln("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	res := 1.0
	var b float64
	for {
		b, res = res, res-(res*res-x)/(2*res)
		if math.Abs(res-b) < 1e-6 {
			return res, nil
		}
	}
	return res, nil
}

func main() {
	fmt.Println(Sqrt(-2))
}
