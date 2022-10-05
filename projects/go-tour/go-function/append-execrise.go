package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	// 1343
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile("./append-execrise.go")
	b = digitRegexp.Find(b)
	a := append(make([]byte, 0), b...)
	fmt.Println(string(a))
}
