package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println("index:", i, "arg:", arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
