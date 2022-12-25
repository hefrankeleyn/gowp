package main

import "fmt"

type Pet interface {
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
	fmt.Printf("dog 的名称是 %v\n", dog.Name())
	var pet Pet = dog
	dog.SetName("moster")
	fmt.Printf("dog 名称为：%v\n", dog.Name())
	fmt.Printf("pet 是一个:%v, 它的名字是：%v\n", pet.Category(), pet.Name())
	fmt.Println()
	// 示例二；
	dog1 := Dog{"little pig"}
	fmt.Printf("第一个dog的名称是：%q\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("第二个dog的名称是：%q\n", dog2.Name())
	dog1.name = "moster"
	fmt.Printf("第一个dog的名称是：%q\n", dog1.Name())
	fmt.Printf("第二个dog的名称是：%q\n", dog2.Name())
	fmt.Println()

}
