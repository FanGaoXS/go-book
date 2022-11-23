package main

import (
	"fmt"
	"go-book/ch3/basename/basename1"
	"go-book/ch3/basename/basename2"
)

func main() {
	filename := "a/b.c.go"
	fmt.Println(basename1.Basename(filename))
	fmt.Println(basename2.Basename(filename))
}
