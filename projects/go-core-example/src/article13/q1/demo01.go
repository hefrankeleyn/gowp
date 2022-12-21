package main

import "fmt"

type AnimalCategory struct {
	c01 string
	c02 string
}

func (a AnimalCategory) String() string {
	return fmt.Sprintf("%s, %s", a.c01, a.c02)
}

type Animal struct {
	name string
	AnimalCategory
}

// func (a Animal) String() string {
// 	return fmt.Sprintf("%s,%s", a.name, a.AnimalCategory)
// }

func main() {
	ac := AnimalCategory{c01: "aa", c02: "bb"}
	one_a := Animal{name: "1", AnimalCategory: ac}
	fmt.Printf("%s, %s\n", one_a, one_a.c01)
}
