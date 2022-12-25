package main

import (
	"fmt"
	"reflect"
)

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
	var dog1 *Dog
	fmt.Println("dog1 是 nil")
	dog2 := dog1
	if dog2 == nil {
		fmt.Println("dog2 是 nil")
	} else {
		fmt.Println("dog2 不是 nil")
	}

	var pet Pet = dog2
	if pet == nil {
		fmt.Println("pet 是nil")
	} else {
		fmt.Println("pet 不是 nil")
	}

	fmt.Printf("pet 的类型是 %T\n", pet)
	fmt.Printf("pet 的类型是 %s\n", reflect.TypeOf(pet).String())
	fmt.Printf("dog1 的类型是  %T\n", dog1)
	fmt.Printf("dog2 的类型是  %T\n", dog2)
	fmt.Println()

	// 示例二：
	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("pet 是nil")
	} else {
		fmt.Println("pet 不是 nil")
	}
}
