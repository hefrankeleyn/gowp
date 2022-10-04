package main

import "fmt"

const (
	// 通过把1字节向左移动100位，创建一个大数
	// 换句话说，这个二进制数是1后面跟100个0
	Big = 1 << 100
	// 再次将它向右移动99位，因此，我们最终得到 1<<1  或者 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// 会报错
	fmt.Println(needInt(Big))
}
