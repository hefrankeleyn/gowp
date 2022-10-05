package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := range res {
		a := make([]uint8, dx)
		for j := range a {
			a[j] = uint8((i + j) / 2)
		}
		res[i] = a
	}
	return res
}

func main() {
	pic.Show(Pic)
}
