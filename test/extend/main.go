package main

import "fmt"

func main() {
	c := C{A: Aa{}}
	c.MethodOfA()
	c.MethodOfB()
}

type C struct {
	A
	B
}

type A interface {
	MethodOfA()
}

type Aa struct{}

func (Aa) MethodOfA() { fmt.Println("i am method of A") }

type B struct{}

func (B) MethodOfB() { fmt.Println("i am method of B") }
