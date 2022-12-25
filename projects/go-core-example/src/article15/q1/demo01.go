package main

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}
