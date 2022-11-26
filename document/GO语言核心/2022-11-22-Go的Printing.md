# Go的Printing

[toc]

## 一、动作

[fmt#Printing](https://pkg.go.dev/fmt#hdr-Printing)

### 1.1 一般

- `%v` 默认格式的值
- `%+v`  多一个加号标识，能打印struct类型的属性

```go
package main

import (
	"fmt"
)

type School struct {
	name string
	s_id int
}

func main() {
	example := "hello, world"
	example1 := School{name: "s01", s_id: 12}
	// %v 示例1：hello, world， 示例二:{s01 12}
	fmt.Printf("%%v 示例1：%v， 示例二:%v\n", example, example1)
	// %+v 示例1：hello, world， 示例二:{name:s01 s_id:12}
	fmt.Printf("%%+v 示例1：%+v， 示例二:%+v\n", example, example1)
}
```

- `%T` 值的类型

  ```go
  func main() {
  	example := "hello, world"
  	example1 := School{name: "s01", s_id: 12}
  	// %T, 示例1：string, 示例二：main.School
  	fmt.Printf("%%T, 示例1：%T, 示例二：%T\n", example, example1)
  }
  ```

- `%%` 一个百分号的字面量

  ```go
  	// %%
  	fmt.Printf("百分号的字面量，不会消耗值%%v\n")
  ```

### 1.2 布尔类型

- `%t` 一个为true或false的单词

  ```go
    example2 := true
  	// %t 一个为ture或false的单词
  	fmt.Printf("%t\n", example2)
  ```

### 1.3 整数

- `%b` 二进制

  ```go
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	b01 := 8
  	// 二进制（%b）：1000
  	fmt.Printf("二进制（%%b）：%b\n", b01)
  }
  ```

- `%c`  由相应的 Unicode 代码点表示的字符

  ```go
  	c01 := 65
  	// A
  	fmt.Printf("%c\n", c01)
  ```

- `%d` 十进制

  ```go
  	c01 := 65
  	// Unicode代码点所表示的字符（%c）：A
  	fmt.Printf("Unicode代码点所表示的字符（%%c）：%c\n", c01)
  	// 十进制（%d）：65
  	fmt.Printf("十进制（%%d）：%d\n", c01)
  ```

- `%o`  八进制

  ```go
  	c01 := 65
  	// 八进制（%o）：101
  	fmt.Printf("八进制（%%o）：%o\n", c01)
  ```

- `%O` 八进制，并带有`0o`前缀

  ```go
  	c01 := 65
    // 八进制（%O）：0o101
  	fmt.Printf("八进制（%%O）：%O\n", c01)
  ```

- `%q`  使用单引号扩起来的字符字面量，必要时，采用安全的转义表示

  ```go
  	c01, c02 := 65, 1
  	// Unicode代码点所表示的字符（%c）：A
  	// Unicode代码点所表示的字符（%c）：c01: A, c02:
  	fmt.Printf("Unicode代码点所表示的字符（%%c）：c01: %c, c02: %c\n", c01, c02)
  	// 使用单引号扩起来的字符字面量，必要时，采用安全的转义表示
  	// 使用单引号扩起来的字符字面量（%q）：c01: 'A', c02: '\x01'
  	fmt.Printf("使用单引号扩起来的字符字面量（%%q）：c01: %q, c02: %q\n", c01, c02)
  ```

- `%x`  基于16进制，使用小写字母`a-f`标识

  ```go
  	c03 := 31
  	// 十六进制，使用a-f表示（%x）：1f
  	fmt.Printf("十六进制，使用a-f表示（%%x）：%x\n", c03)
  ```

- `%X`  十六进制，使用A-F表示

  ```go
  	c03 := 31
  	// 十六进制，使用A-F表示（%X）：1F
  	fmt.Printf("十六进制，使用A-F表示（%%X）：%X\n", c03)
  ```

- `%U` Unicode 格式

  ```go
  	b01 := 8
    b02 := 1234
  	// Unicode格式（%U）：b01：U+0008， b02：U+04D2
  	fmt.Printf("Unicode格式（%%U）：b01：%U， b02：%U\n", b01, b02)
  ```

### 1.4 浮点类型和复数

- `%b`  无小数部分、二进制指数的科学记数法

- `%e` 科学计数法
- `%E` 科学记数法
- `%f`  浮点数，但无指数部分
- `%F` 等价于`%f`

- `%g`  根据情况使用`%e`或`%f`
- `%G`  根据情况使用`%E`或`%F`
- `%x`  十六进制表示法（具有两个指数的十进制幂）
- `%X`  大写十六进制表示法

```go
func main() {
	v01, v03, v02 := 40.0, 400000000.0, 12.32
	// 二进制指数的科学记数法(%b)：5629499534213120p-47
	fmt.Printf("二进制指数的科学记数法（%%b）%b\n", v01)
	// 科学计数法（%e）：4.000000e+01
	fmt.Printf("科学计数法（%%e）：%e\n", v01)
	// 科学计数法（%E)：4.000000E+01
	fmt.Printf("科学计数法（%%E)：%E\n", v01)
	// 浮点数，但是没有小数部分（%f)：40.000000
	fmt.Printf("浮点数，但是没有小数部分（%%f)：%f\n", v01)
	// 等价于%f（%F)：40.000000
	fmt.Printf("等价于%%f（%%F)：%F\n", v01)
	// 根据情况使用%e或%f（%g）：4e+08, 12.32
	fmt.Printf("根据情况使用%%e或%%f（%%g）：%g, %g\n", v03, v02)
	// 根据情况使用%E或%F（%G）：4E+08, 12.32
	fmt.Printf("根据情况使用%%E或%%F（%%G）：%G, %G\n", v03, v02)
	// 十六进制表示法（具有两个指数的十进制幂）（%x）：0x1.7d784p+28, 0x1.8a3d70a3d70a4p+03
	fmt.Printf("十六进制表示法（具有两个指数的十进制幂）（%%x）：%x, %x\n", v03, v02)
	// 大写十六进制表示法（%X）：0X1.7D784P+28, 0X1.8A3D70A3D70A4P+03
	fmt.Printf("大写十六进制表示法（%%X）：%X, %X\n", v03, v02)
}
```

### 1.5 字符串和切片类型

- `%s`  字符串或byte类型的切片

  ```go
  	v01, v02 := "hello", []byte{'1', '2', 'a'}
  	// 字符串或byte类型的切片（%s）：hello, 12a
  	fmt.Printf("字符串或byte类型的切片（%%s）：%s, %s\n", v01, v02)
  ```

- `%q`  使用 Go 语法安全转义的双引号字符串

  ```go
  	v01, v02, v03 := "hello", []byte{'1', '2', 'a'}, "\"aa"
  	// 字符串中带有转译符号，使用%q的时候，转义符号也能够保留
  	// 用 Go 语法安全转义的双引号字符串（%q）："hello", "12a", "\"aa"
  	fmt.Printf("使用 Go 语法安全转义的双引号字符串（%%q）：%q, %q, %q\n", v01, v02, v03)
  ```

- `%x`  小写十六进制表示

  ```go
  	v01, v02, v03 := "hello", []byte{'1', '2', 'a'}, "\"aa"
  	// 大写十六进制表示（%X）68656C6C6F, 313261
  	fmt.Printf("大写十六进制表示（%%X）%X, %X\n", v01, v02)
  ```

- `%X`  大写十六进制表示

  ```go
  	v01, v02, v03 := "hello", []byte{'1', '2', 'a'}, "\"aa"
  	// 大写十六进制表示（%X）68656C6C6F, 313261
  	fmt.Printf("大写十六进制表示（%%X）%X, %X\n", v01, v02)
  ```

### 1.6 切片地址

- `%p` 以 16 进制表示法表示的第 0 个元素的地址，前导 0x

  ```go
    v02, v03 := []byte{'1', '2', 'a'}
    // 以 16 进制表示法表示的第 0 个元素的地址，前导 0x（%p）：0xc0000b2002
  	fmt.Printf("以 16 进制表示法表示的第 0 个元素的地址，前导 0x（%%p）：%p\n", v02)
  ```

### 1.7 指针

- `%p` 16 进制表示法，前导 0x。%b、%d、%o、%x 和 %X 动词也适用于指针，将值完全格式化为整数。

```go
	p01 := &v01
	// 指针（%p）：0xc000096210, 824634335760
	fmt.Printf("指针（%%p）：%p, %d\n", p01, p01)
```

### 1.8 `%v` 的默认值

- 对于bool类型，默认使用`%t`；

- （int、int8等）整数，默认使用`%d`；
- （uint、uint8等）整数，默认默认使用`%d`；
- flloat32、complex64等，使用：`%g`；
- string：使用`%s`
- chan： 使用`%p`
- pointer 使用：`%p`

### 1.9 复合对象的打印规则

对于复合对象， 元素使用下面的规则：

```go
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
指向上面的对象的指针:   &{}, &[], &map[]
```

```go
	m01 := map[int]string{1: "aa", 2: "bb"}
	mp01 := &m01
	// &map[1:aa 2:bb]
	fmt.Printf("%v\n", mp01)
```

### 1.10 宽度和精度

宽度由动词前的可选十进制数指定。如果不存在，宽度是表示值所必须的。

精度在（可选的）宽度后由一个句点和一个十进制数指定。如果不存在句号，就是用默认精度，后面没有数字的句点指定精度为零。

```go
func main() {
	v01, v02 := 132.1, 0.2
	fmt.Printf("%6.1f\n", v01)
	fmt.Printf("%6.1f\n", v02)
}
```

```shell
$ go run demo06_width.go
 132.1
   0.2
$
```

`%f` 默认宽度，默认精度

`%9f`  宽度是9， 默认精度

`%.2f`  默认宽度，精度为2

`%9.2f`  宽度是， 精度是2

`%9.f`   宽度是9，精度是0

### 1.11 其它的符号

- `+`  总是为数值打印一个符号；

  ```go
  	d01 := 2
  	// +2
  	fmt.Printf("%+d\n", d01)
  ```

- `-` 在右边而不是左边填充空格

  ```go
    v02 := 0.2
  	//   0.2
  	fmt.Printf("%6.1f\n", v02)
  	// 0.2
  	fmt.Printf("%-6.1f\n", v02)
  ```

- `#`  备用格式：为二进制 (%#b) 添加前导 0b，为八进制 (%#o) 添加 0，

  0x 或 0X 十六进制（%#x 或 %#X）； 为 %p (%#p) 去掉 0x；

- ` `（空格）对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格

- `0` 用前导零而不是空格填充；
  对于数字，这会在符号后移动填充；
  忽略字符串、字节切片和字节数组

  ```go
  	v01, v02 := 132.1, 0.2
  	fmt.Printf("%6.1f\n", v01)
  	fmt.Printf("%06.1f\n", v01)
  ```

  ```shell
  $ go run demo06_width.go
   132.1
  0132.1
  $
  ```

  





