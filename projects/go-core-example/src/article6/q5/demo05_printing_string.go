package main

import (
	"fmt"
)

func main() {
	v01, v02, v03 := "hello", []byte{'1', '2', 'a'}, "\"aa"
	// 字符串或byte类型的切片（%s）：hello, 12a
	fmt.Printf("字符串或byte类型的切片（%%s）：%s, %s\n", v01, v02)
	// 字符串中带有转译符号，使用%q的时候，转义符号也能够保留
	// 用 Go 语法安全转义的双引号字符串（%q）："hello", "12a", "\"aa"
	fmt.Printf("使用 Go 语法安全转义的双引号字符串（%%q）：%q, %q, %q\n", v01, v02, v03)
	// 小写十六进制表示（%x）68656c6c6f, 313261
	fmt.Printf("小写十六进制表示（%%x）%x, %x\n", v01, v02)
	// 大写十六进制表示（%X）68656C6C6F, 313261
	fmt.Printf("大写十六进制表示（%%X）%X, %X\n", v01, v02)
	// 以 16 进制表示法表示的第 0 个元素的地址，前导 0x（%p）：0xc0000b2002
	fmt.Printf("以 16 进制表示法表示的第 0 个元素的地址，前导 0x（%%p）：%p\n", v02)
	p01 := &v01
	// 指针（%p）：0xc000096210, 824634335760
	fmt.Printf("指针（%%p）：%p, %d\n", p01, p01)
	m01 := map[int]string{1: "aa", 2: "bb"}
	mp01 := &m01
	// &map[1:aa 2:bb]
	fmt.Printf("%v\n", mp01)
}
