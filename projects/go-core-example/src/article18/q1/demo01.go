package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	// 卫述语句
	if request == "" {
		// 基本的生成错误值的方法
		// 该值的静态类型为 error，动态类型是包级私有的类型*errorString
		err = errors.New("empty request")
		err2 := fmt.Errorf("empty request")
		_ = err2
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	for _, request := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", request)
		resp, err := echo(request)
		// 卫述语句
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s \n", resp)
	}
}
