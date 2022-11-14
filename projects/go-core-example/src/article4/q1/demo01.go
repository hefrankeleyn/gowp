package main

import (
	"flag"
	"fmt"
)

func main() {
	// 变量的第一种声明方式
	// var name string
	// flag.StringVar(&name, "name", "everyone", "The greeting Object")

	// 变量的第二种声明方式, 返回的类型 为 *string
	// var name = flag.String("name", "everyone", "The greeting Object.")

	// 变量的第三种声明方式
	name := flag.String("name", "everyone", "The greeting Object.")
	flag.Parse()
	fmt.Println(*name)
	fmt.Printf("%T\n", name)
}
