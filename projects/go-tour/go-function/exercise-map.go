package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WorldCount(s string) map[string]int {
	sp := strings.Split(s, " ")
	res := make(map[string]int)
	for _, v := range sp {
		c := res[v]
		c += 1
		res[v] = c
	}
	return res
}

func main() {
	wc.Test(WorldCount)
}
