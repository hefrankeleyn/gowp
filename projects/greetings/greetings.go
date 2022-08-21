package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为一个命名的人返回一个问候
func Hello(name string) (string, error) {
	// 如果没有给出name，返回一个带有消息的错误
	if name == "" {
		return "", errors.New("empty name")
	}
	// 如果接受到一个名字，返回一个问候，嵌套名字在消息中
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// 使用随机消息格式创建消息
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellow 返回一个关联了每一个人名和问候语的map
func Hellos(names []string) (map[string]string, error) {
	// 一个map，用于关联人名和消息
	messages := make(map[string]string)
	// 循环查看接口的分片names，调用Hello函数 为每一个name获取一个消息
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 在map中，关联检索到的消息和name
		messages[name] = message
	}
	return messages, nil
}

// 为在函数中使用的变量初始化集合的最初始的值
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 返回一个问候语集合中的一个，被返回的消息是随机选择的
func randomFormat() string {
	// 一个消息格式的切片
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 根据指定的随机索引从格式化切片中返回一个随机选择的消息格式
	return formats[rand.Intn(len(formats))]
}
