package main

import (
	"fmt"
)

func main() {
	fmt.Println("max of 8bit is", MaxBit(8))
	fmt.Println("max of unsigned 8bit is", MaxUBit(8))
}

func MaxBit(x int64) int64 {
	return Pow(2, x-1+1-1) - 1
}

func MaxUBit(x int64) int64 {
	// unsigned
	return Pow(2, x+1-1) - 1
}

func Pow(base, exponent int64) int64 {
	return base << (exponent - 1) // x^n <=> x<<(n-1)
}
