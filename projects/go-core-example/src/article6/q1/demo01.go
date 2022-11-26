package main

import (
	"fmt"
)

var container = []string{"aa", "bb", "cc"}

func main() {
	container := map[int]string{0: "aa", 1: "bb", 2: "cc"}
	// 第一种判断类型的方法：
	// 语法 x.(T) ， 这里的x必须是接口类型，具体是哪个接口无所谓。
	// interface{} 代表，空接口，任何类型都能很方便地转换成空接口。
	// interface{}(container) 把 container 转为空接口类型
	// .([]string) 判断前面的类型是否为切片类型
	// 这里的ok也可以没有，如果没有，判断不通过会引发异常
	// value 是类型转换之后都值；ok 是断言是否成功
	vlaue, ok := interface{}(container).([]string)
	fmt.Printf("%q, %%, %v\n", vlaue, ok)
	fmt.Printf("这个元素是：%q\n", container)

	// 第二种判断类型的方法
	val, err := getElement(container)
	fmt.Printf("%q, %v", val, err)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("不支持的类型： %T", containerI)
		return
	}
	return
}
