package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashMethod int

func main() {
	flag.Parse()

	var s string
	fmt.Scanln(&s)
	switch hashMethod {
	case 256:
		fmt.Printf("256 = %s", sha256.Sum256([]byte(s)))
	case 384:
		fmt.Printf("384 = %s", sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("512 = %s", sha512.Sum512([]byte(s)))
	}
}

func init() {
	flag.IntVar(&hashMethod, "m", 256, "hash method")
}
