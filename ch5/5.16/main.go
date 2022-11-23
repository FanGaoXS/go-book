package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("str =", Join(", ", "Hello", "World"))
}

func Join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}
