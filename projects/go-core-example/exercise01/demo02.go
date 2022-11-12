package main

import (
	"flag" // go 语言标准库中的一个代码包
	"fmt"
	"os"
)

var name string

// 创建一个私有的命令容器
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	fmt.Println("运行init函数....")
	// 自定义源码文件的参数说明 方案一：
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// 自定义源码文件的参数说明 方案二：
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag 接收四个参数
	//     第一个参数： 用于存储该命令参数值的地址；
	//     第二个参数： 用于指定该命令参数的命令；
	//     第三个参数： 指定在未追加该命令参数时的默认值，这里是everyone；
	//     第四个参数： 该命令参数的简要说明
	// flag.String()  会直接返回一个已经分配好的用于存储命令参数值的地址。
	// var name = flag.String("name", "everyone", "The greeting object.")
	// flag.StringVar(&name, "name", "everyone", "The greeting object.")

	// 使用自定义私有的命令容器
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")

}

func main() {

	fmt.Println("运行main函数.....")
	// 用于真正解析命令参数，并把它们的值赋给相应的变量
	// 该函数的调用必须放在所有命令参数存储载体的声明（这里是对变量name的声明）和设置（这里是falg.StringVar函数的调度）之后，
	// 并且在读取任何命令参数值之前进行。
	// flag.Parse()
	// 使用自定义私有的命令容器
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
