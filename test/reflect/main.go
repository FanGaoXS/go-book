package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var i int = 3
	t := reflect.TypeOf(i)  // reflect.Type
	v := reflect.ValueOf(i) // reflect.Value
	fmt.Printf("type of %v is %v\n", i, t)
	fmt.Printf("value of %v is %v\n", i, v)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))  // *os.File
	fmt.Println(reflect.ValueOf(w)) // &{0xc00005c0c0}
	fmt.Printf("%T\n", w)           // *os.File
	fmt.Printf("%v\n", w)           // *os.File
}
