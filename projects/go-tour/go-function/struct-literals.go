package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 有一个类型Vertex
	v2 = Vertex{X: 1}  // Y 是0 是隐含的
	v3 = Vertex{}      // X 是0， Y是0
	p  = &Vertex{1, 2} // 有一个类型 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
