package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"aa": {51.2, -161.1},
	"bb": {-91.3, 9.91},
}

func main() {
	fmt.Println(m)
}
