package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"aa": Vertex{
		12.1, -9.1,
	},
	"bb": Vertex{
		51.1, -20.1,
	},
}

func main() {
	fmt.Println(m)
}
