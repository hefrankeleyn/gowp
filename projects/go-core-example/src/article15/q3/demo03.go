package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func main() {
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	fmt.Println(dogPtr)

	namePtr := dogPtr + unsafe.Offsetof(dog.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(nameP)
}
