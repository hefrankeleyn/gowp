package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
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

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例一
	dog := Dog{"little dog"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("dog 实现了 Pet类型： %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*dog 实现了 Pet类型：%v\n", ok)
	fmt.Println()

	// 示例二：
	var pet Pet = &dog
	fmt.Printf("这个pet是%s, 名字是：%v\n", pet.Category(), pet.Name())
}
