package main

import (
	. "article3/q4/hw/lib" // 前面加点，可以直接调用里面的函数
	// q4_lib1 "article3/q4/lib"
	_ "article3/q4/lib" // 如果只想引入包，而实际没有调用
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	Hello(name) // 直接调用函数名，不用加 XXX. 了
	// q4_lib1.Hello(name)
	// hw_lib1.Hello(name)
	// in.Hello(os.Stdout, name)
}
