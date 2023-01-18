package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	return s.Name
}
func (s *Student) GetAge() int {
	return s.Age
}

func main() {
	stu := &Student{
		Name: "test",
		Age:  18,
	}
	// methods
	typ := reflect.TypeOf(stu)
	for i := 0; i < typ.NumMethod(); i++ {
		fmt.Println(typ.Method(i))
		//fmt.Println(typ.Method(i).Type)
		//fmt.Println(typ.Method(i).Name)
		//fmt.Println(typ.Method(i).PkgPath)
		//fmt.Println(typ.Method(i).Func)
	}
}
