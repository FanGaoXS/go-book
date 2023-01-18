package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func main() {
	stu := &Student{
		Name: "test",
		Age:  18,
	}
	// tags
	typ := reflect.TypeOf(stu).Elem()
	for i := 0; i < typ.NumField(); i++ {
		fmt.Println(typ.Field(i).Tag)
		//fmt.Println(typ.Field(i).Tag.Get("json"))
	}
}
