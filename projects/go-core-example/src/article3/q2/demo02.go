package main

import (
	// lib2 "article3/q2/lib"
	"article3/q2/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	// hello(name)
	// lib2.Hello(name)
	lib.Hello(name)
}
