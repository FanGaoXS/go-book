package main

import "fmt"

func main() {
	i := 0
	incr(&i)
	fmt.Println("i =", i)
}

func incr(i *int) int {
	*i++
	return *i
}
