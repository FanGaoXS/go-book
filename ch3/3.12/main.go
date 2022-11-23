package main

import (
	"fmt"

	"go-book/ch3/3.12/valid"
)

func main() {
	s1 := "dasds321adsa"
	s2 := "321dsadasdas"
	fmt.Println(valid.IsValid(s1, s2))
}
