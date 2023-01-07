package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting Object")
}

func main() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(greeting, introduce())
}

func introduce() string {
	return "Welcome to my Golang column."
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	} else {
		return fmt.Sprintf("hello, %s!", name), nil
	}
}
