package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() { fmt.Println(i) }()
	}

	for j := 0; j < 100_0000; j++ {
		_ = j
	}
}
