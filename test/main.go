package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello, world!"
	fmt.Println("len(s) =", len(s))
	fmt.Printf("s[0] = %d, s[7] = %d\n", s[0], s[7])

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

func add1(r rune) rune { return r + 1 }
