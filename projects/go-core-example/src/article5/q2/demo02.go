package main

import "fmt"

var container = []string{"one", "two", "three"}

func main() {
	container := map[int]string{0: "one", 2: "two", 3: "three"}
	fmt.Println(container)
}
