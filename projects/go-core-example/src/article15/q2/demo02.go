package main

import "fmt"

type Named interface {
	Name() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func New(name string) Dog {
	return Dog{name}
}

func main() {
	// 常量值总是会被存储到一个确切的内存区域中，并且这种值肯定是不可变的。
	const num = 123
	// _ = &num  // 寻址不可取
	// _ = &(123) // 基本类型的字面量不可寻址

	var str = "abc"
	_ = str
	// _ = &(str[0]) // 对字符串变量的索引结果值不可寻址
	// _ = &(str[0:1]) // 对字符串变量的切片结果值不可寻址
	str02 := str[2]
	_ = &str02 // 但这样的寻址就是合法的

	// _ = &(123 + 456) // 算术操作的结果值不可寻址
	num2 := 456
	_ = num2
	// _ = &(num + num2) // 算术操作的结果值不可寻址

	// _ = &([3]int{1, 2, 3}[1]) // 对数组字面量的索引结果值不可寻址
	// _ = &([3]int{1, 2, 3}[0:1]) // 对数组字面的切片结果值不可寻址
	_ = &([]int{1, 2, 3}[0]) // 对切片的字面量的索引结果却是可寻址的
	// _= &([]int{1, 2,3}[0:2]) // 对切字面量的切片结果值不可寻址
	// _ = &(map[int]string{0:"aa"}[0]) // 对字面字面量的索引结果值不可寻址

	var map01 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map01
	// _ = &(map01[1]) // 对字典变量的索引结果值不可寻址

	// _ = &(func(x, y int) int {
	// 	return x + y
	// })   // 字面量代表的函数不可寻址

	// _ = &(fmt.Sprint) // 标识符代表的函数不可寻址

	// _ = &(fmt.Sprintln("aa")) // 对函数调用的结果值不可寻址

	dog := Dog{"little pig"}
	_ = dog
	// _ = &(dog.Name) // 标识符代表的函数不可寻址
	// _ = &(dog.Name())  // 对方法的调用结果值不可寻址

	// _ = &(Dog{"little pig"}.name) // 对结构体字面量的字段不可寻址

	// _ = &(interface{}(dog)) // 类型转换表达式的结果值不可寻址
	dogI := interface{}(dog)
	_ = dogI
	// _ = &(dogI.(Named))  // 类型断言表达式的结果值不可寻址
	named := dogI.(Named)
	_ = named
	// _ = &(named.(Dog)) // 类型断言表达式的结果值不可寻址

	var ch chan int = make(chan int, 1)
	ch <- 1
	// _ = &(<- ch) //  接收表达式的结果值不可寻址

	p := New("Big pig")
	p.SetName("mon")
	// New("Big pig").SetName("mon") // 不可寻址，链式调用出错

	map[string]int{"a": 2}["a"]++
	var m map[string]int = map[string]int{"b": 2}
	m["b"]++

	var m01 map[string]int = map[string]int{"a": 2}
	m01["a"] = 3

	for mk := range m01 {
		fmt.Println(m01[mk])
	}

}
