package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	if m == nil {
		fmt.Println(m)
		fmt.Println("m is nil")
	}
	m = make(map[string]Vertex)
	m["aa"] = Vertex{23.1, -12.1}
	fmt.Println(m["aa"])
}
