package greetings

import "fmt"

// Hello 为一个命名的人返回一个问候
func Hello(name string) string {
	// 返回一个问候，嵌套名字在消息中
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
