# unicode和字符编码

[toc]

## 一、Go语言字符编码基础

Go语言的所有源代码，都必须按照Unicode编码规范中的UTF-8编码格式进行编码。

> 换句话说，Go语言的源码文件必须使用UTF-8编码格式进行编码。

## 二、ASCII编码和Unicode编码

ASCII编码使用一个字节，只对拉丁字母进行编码。

> ASCII 码一共规定了128个字符的编码。英语用128个符号编码就够了。
>
> 其他语言，128个符号是不够的。
>
> 在所有用ASCII码表示的语言中，0--127表示的符号是一样的，128--255的这一段不同语言表示的符号不一样。

Unicode编码是以ASCII编码集为出发点，实现了一种更加通用的、针对书面字符和文本的字符标准。

> Unicode 编码为世界上现存的所有自然语言中的每一个字符，都设定了一个唯一的二进制编码。
>
> 它定义了不同自然语言的文本数据在国际间交换的统一方式，并为全球化软件创建了一个重要的基础。

Unicode 只是一个符号集，它只规定了符号的二进制代码，却没有规定这个二进制代码应该如何存储。

UTF-8 是 Unicode 的实现方式之一。

UTF-8 最大的一个特点，就是它是一种变长的编码方式。它可以使用1~4个字节表示一个符号，根据不同的符号而变化字节长度。

参考：[字符编码笔记：ASCII，Unicode 和 UTF-8](http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)。

## 三、一个string类型的值在底层的表达方式

在底层，一个string类型的值是由一系列相对应的Unicode代码点点UTF-8编码值表示的。

在Go语言中，一个string类型的值即可以被拆分成一个包含多个**字符**的序列，也可以拆分成一个包含多个**字节**的序列。

字符序列用rune为元素类型的切片来表示，字节是用byte元素类型的切片表示。

rune是Go语言特有的一个基本数据类型，它的一个值代表一个字符，即：一个Unicode字符。

```go
type rune = int32
```

rune 实际上是一个int32的别名类型，它用四个字节的存储空间，总是能存下一个UTF-8编码值。

```go
	var str string
	str = "Go 爱好者"
	//The string："Go 爱好者"
	fmt.Printf("The string：%q\n", str)
	// runes(char)： ['G' 'o' ' ' '爱' '好' '者']
	fmt.Printf("runes(char)： %q\n", []rune(str))
	// runes(char)： ['G' 'o' ' ' '爱' '好' '者']
	fmt.Printf("runes(char)： %q\n", []rune(str))
	//runes(byte)： [47 6f 20 e7 88 b1 e5 a5 bd e8 80 85]
	fmt.Printf("runes(byte)： [% x]\n", []byte(str))
```

## 四、使用range遍历字符串

带有range子句的for语句会把被遍历的字符串值拆成一个**字节序列**，然后试图找出这个字节序列中包含的每一个UTF-8编码值，或者说每一个Unicode字符。

这样的for语句可以为两个迭代变量赋值。如果存在两个迭代变量：第一个变量的值，是当前字节序列中的某个UTF-8编码值的第一个字节对应的索引值；第二个变量，就是这个UTF-8编码值代表的那个Unicode字符，其类型会是rune。

```go
	var str string
	str = "Go 爱好者"
	for byte_index, char_val := range str {
		fmt.Printf("%d : %q [% x]\n", byte_index, char_val, []byte(string(char_val)))
	}
	/*
		0 : 'G' [47]
		1 : 'o' [6f]
		2 : ' ' [20]
		3 : '爱' [e7 88 b1]
		6 : '好' [e5 a5 bd]
		9 : '者' [e8 80 85]
	*/
```

如果这样的for语句只赋值一个变量值，这个变量值是：是当前字节序列中的某个UTF-8编码值的第一个字节对应的索引值。





