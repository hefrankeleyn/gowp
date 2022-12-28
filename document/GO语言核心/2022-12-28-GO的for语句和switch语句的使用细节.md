# GO的for语句和switch语句的使用细节

[toc]

## 一、携带range子句的for语句使用细节

### （1）range 左边只有一个迭代变量，该迭代变量代表索引值

```go
func main() {
	number1 := []int{1, 2, 3, 4, 5, 6, 7}
	for i := range number1 {
		if i == 3 {
			number1[i] |= i
		}
	}
	fmt.Println(number1)
}
```

> 上面语句打印：[1 2 3 7 5 6 7]

### （2）range 迭代数组，迭代的是数据的副本

```go
func main() {
	number2 := [...]int{1, 2, 3, 4, 5}
	maxIndex := len(number2) - 1
	for i, e := range number2 {
		if i == maxIndex {
			number2[0] += e
		} else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
}
```

> 上面语句打印：[6 3 5 7 9]

- range表达式只会在for语句开始执行时被求值一次，无论后面会有多少次迭代；
- range表达式的求值结果会被复制，也就是说，被迭代对象是range表达式结果值的副本而不是原值；

```go
func main() {
	// number2 := [...]int{1, 2, 3, 4, 5}
	number2 := []int{1, 2, 3, 4, 5}
	maxIndex := len(number2) - 1
	for i, e := range number2 {
		if i == maxIndex {
			number2[0] += e
		} else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
}
```

> 上面打印：[16 3 6 10 15]

## 二、switch语句中，switch表达式与case表达式之间的联系

### （1）case子句的选择

一旦某个case子句被选中，其中的附带在case表达式后面的那些语句就会被执行。与此同时，其它的所有case子句都会被忽略。

如果被选中的case子句附带的语句列表中包含了fallthrough语句，那么紧挨在它下边的那个case子句附带的语句也会被执行。

```go
func main() {
	num := 3
	switch num {
	case 3:
		fmt.Println("3 case")
		fallthrough
	case 2:
		fmt.Println("2 case")
	case 1:
		fmt.Println("1 case")
	default:
		fmt.Println("no case")
	}
}
```

> 上面的语句打印：
>
> 3 case
> 2 case

### （2）switch表达式的结果值，默认的类型转换

如果switch表达式的结果值是无类型的常量，那么这个常量会被自动地转换为此种常量的默认类型的值。

```go
func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 1 {
	case value1[0], value[1]:
		fmt.Println("0 or 1")
	default:
		fmt.Println("other")
	}
}
```

上面的代码会变异报错，因为switch表达式的 1+1 会被转化为 int类型，和case表达式的类型不匹配。

### （3）case表达式的结果值，默认的类型转换

如果case表达式中子表达式的结果值是无类型的常量，那么它的类型会被自动地转换为switch表达式的结果类型。

```go
func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case 0, 1:
		fmt.Println("0 or 1")
	default:
		fmt.Println("other")
	}
}
```

上面的代表可以编译通过。

### （4）switch语句对case表达式的约束

switch语句不允许case表达式中的子表达式结果值存在相等的情况，不论这鞋结果值相等的子表达式，是否存在于不同的case表达式中。

**这个约束只针对结果值为常量的子表达式。**

```go
// 编译失败
func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 1, 2:
		fmt.Println("1 or 2")
	default:
		fmt.Println("other")
	}
}
```

```go
// 编译失败
func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case 0, 1, 1:
		fmt.Println("0 or 1")
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	default:
		fmt.Println("other")
	}
}
```

```go
// 编译成功
func main() {
	value1 := []int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[1], value1[2]:
		fmt.Println("1 or 2")
	default:
		fmt.Println("other")
	}
}
```

这种绕过方式对类型判断的switch语句就无效了。

> 可以用 `.(type)` 断言来获取 `interface{}`的真实类型 。

```go
// 编译报错
func main() {
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint8, uint16:
		fmt.Println("int8 or int16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}
}
```

上面的编译报错，因为byte类型是uint8类型的别名类型。