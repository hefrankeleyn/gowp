package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// 设置Logger预定义的属性，包含了log项的前缀，和一个禁用打印时间，源文件，行号的标识。
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 获取一个问候消息，并打印它
	// message := greetings.Hello("hef")
	// 请求一个问候消息
	message, err := greetings.Hello("")
	// 如果一个错误返回，将它打印到控制台，并退出程序
	if err != nil {
		log.Fatal(err)
	}

	// 如果没有错误返回，打印返回到消息到控制台
	fmt.Println(message)
}
