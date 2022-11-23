package main

import (
	"fmt"
	"go-book/ch2/2.1/tempconv"
)

func main() {
	var c tempconv.Celsius = 36.5
	fmt.Println("c =", c)
	fmt.Println("c2f =", tempconv.CToF(c))
	fmt.Println("c2k =", tempconv.CToK(c))

	var f tempconv.Fahrenheit = 97.7
	fmt.Println("f =", f)
	fmt.Println("f2c =", tempconv.FToC(f))
	fmt.Println("f2k =", tempconv.FToK(f))

	var k tempconv.Kelvin = 0
	fmt.Println("k =", k)
	fmt.Println("k2c =", tempconv.KToC(k))
	fmt.Println("k2f =", tempconv.KToF(k))
}
