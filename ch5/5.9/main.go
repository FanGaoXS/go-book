package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "foo2313213"
	res := expand(str, func(s string) string {
		return "abc"
	})
	fmt.Println("res =", res)
}

func expand(s string, f func(string) string) string {
	res := f(s)

	return strings.ReplaceAll(s, "foo", res)
}
