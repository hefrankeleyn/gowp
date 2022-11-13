package main

import (
	"article3/q3/lib"
	// in "article3/q3/lib/internal"
	"flag"
	// "os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	// in.Hello(os.Stdout, name)
}
