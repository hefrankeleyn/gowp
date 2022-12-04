package main

import (
	"container/list"
	"fmt"
)

func main() {
	oneList := list.New()
	// fmt.Printf("oneList： %T, %v\n", oneList, oneList.Len())
	e1 := oneList.PushBack(4)
	// fmt.Printf("e1： %T, %v\n", e1, e1.Value)
	e2 := oneList.PushFront(1)
	// fmt.Printf("oneList: %v, e2： %T, %v\n", oneList, e2, e2.Value)
	oneList.InsertAfter(5, e1)
	oneList.InsertBefore(8, e2)
	fmt.Printf("链表的长度为：%d\n", oneList.Len())
	// oneList.Init() // 初始化链表，或清空链表
	for e := oneList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
