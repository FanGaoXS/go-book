package main

import "fmt"

func main() {
	myMap := make(map[string]string)
	value, _ := myMap[""], false
	fmt.Println("value =", value)
}
