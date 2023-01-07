package q2

import "math"

// GetPrimes 用于获取所有小于或等于参数max的所有质数
// 本函数使用的是爱拉托逊斯筛选法（Sieve Of Eratosthenes）。
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	makes := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if makes[i] == false {
			for j := i * i; j < max; j += i {
				if makes[j] == false {
					makes[j] = true
					count++
				}
			}
		}
	}
	prims := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if makes[i] == false {
			prims = append(prims, i)
		}
	}
	return prims
}
