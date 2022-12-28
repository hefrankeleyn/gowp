package main

import "fmt"

// byte是uint8 的别名
func main() {
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	// case uint8, uint16:
	// 	fmt.Println("int8 or int16")
	// case byte:
	// 	fmt.Println("byte")
	default:
		fmt.Printf("unsupported type: %T\n", t)
	}
}
