package greetings

import (
	"errors"
	"fmt"
)

// Hello 为一个命名的人返回一个问候
func Hello(name string) (string, error) {
	// 如果没有给出name，返回一个带有消息的错误
	if name == "" {
		return "", errors.New("empty name")
	}
	// 如果接受到一个名字，返回一个问候，嵌套名字在消息中
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
