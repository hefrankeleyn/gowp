package main

import (
	"fmt"
)

type mystring string

func f[P ~float32 | ~float64]() {
	v11 := P(1.1)
	fmt.Printf("%T, %v\n", v11, v11)
}

func main() {
	iota := 2
	v01 := uint(iota)
	v02 := float32(2.718281828)
	v03 := complex128(1)
	fmt.Println(v01)
	fmt.Println(v02)
	fmt.Println(v03)
	v04 := float32(0.49999999)
	fmt.Println(v04)
	v05 := float64(-1e-1000)
	fmt.Println(v05)
	v06 := string('x')
	fmt.Println(v06)
	v07 := string(0x266c)
	fmt.Println(v07)
	v08 := mystring("foor" + "bar")
	fmt.Printf("%T, %v\n", v08, v08)
	v09 := string([]byte{'a'})
	fmt.Printf("%T, %v\n", v09, v09)
	v10 := (*int)(nil)
	fmt.Printf("%T, %v\n", v10, v10)
	// v11 := int(1.2)
	// fmt.Println(v11)
	// v12 := string(1.2)
	fmt.Println(string(-1))

}
