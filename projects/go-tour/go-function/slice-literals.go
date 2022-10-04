package main

import "fmt"

func main() {
	a := [3]int{2, 3, 5}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%T， %v\n", q, q)
	fmt.Printf("%T, %v\n", a, a)

	r := []bool{true, false, true, true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, false},
	}
	fmt.Println(s)
}
