package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i += 1
}

func b() {
	defer fmt.Println()
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello ")
	a()
	b()
	fmt.Println(c())

}
