package main

import (
	"fmt"
)

func main() {
	b01 := 8
	// 二进制（%b）：1000
	fmt.Printf("二进制（%%b）：%b\n", b01)
	c01, c02 := 65, 1
	// Unicode代码点所表示的字符（%c）：A
	// Unicode代码点所表示的字符（%c）：c01: A, c02:
	fmt.Printf("Unicode代码点所表示的字符（%%c）：c01: %c, c02: %c\n", c01, c02)
	// 十进制（%d）：65
	fmt.Printf("十进制（%%d）：%d\n", c01)
	// 八进制（%o）：101
	fmt.Printf("八进制（%%o）：%o\n", c01)
	// 八进制（%O）：0o101
	fmt.Printf("八进制（%%O）：%O\n", c01)
	// 使用单引号扩起来的字符字面量，必要时，采用安全的转义表示
	// 使用单引号扩起来的字符字面量（%q）：c01: 'A', c02: '\x01'
	fmt.Printf("使用单引号扩起来的字符字面量（%%q）：c01: %q, c02: %q\n", c01, c02)
	c03 := 31
	// 十六进制，使用a-f表示（%x）：1f
	fmt.Printf("十六进制，使用a-f表示（%%x）：%x\n", c03)
	// 十六进制，使用A-F表示（%X）：1F
	fmt.Printf("十六进制，使用A-F表示（%%X）：%X\n", c03)
	b02 := 1234
	// Unicode格式（%U）：b01：U+0008， b02：U+04D2
	fmt.Printf("Unicode格式（%%U）：b01：%U， b02：%U\n", b01, b02)
}
