package main

import (
	"fmt"
)

func main() {
	v01, v02 := 132.1, 0.2
	fmt.Printf("%6.1f\n", v01)
	fmt.Printf("%06.1f\n", v01)
	//   0.2
	fmt.Printf("%6.1f\n", v02)
	// 0.2
	fmt.Printf("%-6.1f\n", v02)
	d01 := 2
	// +2
	fmt.Printf("%+d\n", d01)

}
