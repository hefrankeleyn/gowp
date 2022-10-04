package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("go 运行在：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

}