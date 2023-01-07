# GO的panic、recove函数，以及defer语句

[toc]

## 一、运行时恐慌panic

panic是一种在运行时抛出来的异常。比如"index of range"。

panic的详情：

```go
package main

import "fmt"

func main() {
	oneC := []int{1, 2, 3, 4, 5}
	v5 := oneC[5]
	fmt.Println(v5)
}
```

```shell
$ go run demo01.go
panic: runtime error: index out of range [5] with length 5

goroutine 1 [running]:
main.main()
	/Users/lifei/Documents/workspace/githubRepositoies/gowp/projects/go-core-example/src/article19/q1/demo01.go:7 +0x1b
exit status 2
$
```

- 打印信息的第一行，"panic: "右边的内容，正是panic包含的`runtime.Error`类型值的字符串表示形式；

- “goroutine 1 [running]” 表示有一个id为 1 的 goroutine在此panic被引发的时候正在运行；

  > 这里的ID编号并不重要，是GO语言运行时系统内部给予的一个goroutine编号，我们在程序中无法获取，也无法改变。

- 再下面是指出哪一行发生错误。“+0x1b”代表 此行代码相对于其所属函数的入口程序计数偏移量， 一般用途不大。

- 最后的 “exit status 2”，表明我的这个程序是以退出状态码2结束运行的。

  > 在大多操作系统中，只要退出状态码不是0，都意味着程序运行的非正常结束。

## 二、panic被引发到程序终止，经历的过程

某个函数无疑触发了panic：

- 初始的panic详情会被建立起来，此行代码所属函数的执行随机终止。
- 控制权立刻转移到上一级；
- 控制权如此一层层沿着调用栈的反方向传播至顶端，也就是我们编写的最外层函数；
- 最终，控制权被GO语言运行时系统收回。随后程序崩溃并终止运行；

panic 详情会在控制权传播的过程中，被逐渐地积累和完善，并且，控制权会一级一级地沿着调用栈的反方向传播至顶端。因此，在针对某个 goroutine 的代码执行信息中，调用栈底端的信息会先出现，然后是上一级调用的信息，以此类推，最后才是此调用栈顶端的信息。

## 三、有意引发一个panic，让panic包含一个值

- 可以使用panic函数有意地引发一个 panic。
- 在调用panic函数时，把某个值作为参数传给该函数就可以了。由于panic函数的唯一一个参数是空接口（也就是interface{}）类型的，所以从语法上讲，它可以接受任何类型的值。
- 但是，我们最好传入error类型的错误值，或者其他的可以被有效序列化的值。这里的“有效序列化”指的是，可以更易读地去表示形式转换。

打印错误信息：

- 对于fmt包下的各种打印函数来说，error类型值的Error方法与其他类型值的String方法是等价的，它们的唯一结果都是string类型的；
- 如果某个值有可能会被记到日志里，那么就应该为它关联String方法。

## 四、施加应对 panic 的保护措施，从而避免程序崩溃

联用defer语句和recover函数调用，才能够恢复一个已经发生的 panic。

GO语言的内建函数recover专门用于恢复panic。recover函数无需任何参数，并且会返回一个空接口类型的值。

defer 语句用来延迟执行代码。延迟到该语句所在的函数即将执行结束的那一刻，无论结束执行的原因是什么。

> 限制：有一些调用表达式是不能出现在这里的，包括：针对 Go 语言内建函数的调用表达式，以及针对unsafe包中的函数的调用表达式。

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main")
	// 延迟func函数的执行，直到main结束
	defer func() {
		fmt.Println("Enter defer function")
		if p := recover(); p != nil {
			fmt.Printf("%v\n", p)
		}
		fmt.Println("Exit defer function")
	}()
	// 引发painc
	panic(errors.New("soming wrong"))
	fmt.Println("Exit function main")
}
```

## 五、多条defer语句，多条defer语句的执行顺序

在同一个函数中，defer函数调用的执行顺序与它们分别所属的defer语句的出现顺序（更严谨地说，是执行顺序）完全相反。

当一个函数即将结束执行时，其中的写在最下边的defer函数调用会最先执行，其次是写在它上边、与它的距离最近的那个defer函数调用，以此类推，最上边的defer函数调用会最后一个执行。

defer语句执行的内幕：

- 在defer语句每次执行的时候，Go 语言会把它携带的defer函数及其参数值另行存储到一个链表中。

  > 在defer语句每次执行的时候，Go 语言会把它携带的defer函数及其参数值另行存储到一个链表中。

- 在需要执行某个函数中的defer函数调用的时候，Go 语言会先拿到对应的链表，然后从该链表中一个一个地取出defer函数及其参数值，并逐个执行调用。

```go
package main

import "fmt"

func main() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for %d\n", i)
	}
	defer fmt.Println("last defer")
}
```

