package main

import "fmt"

func main() {
	var res []func()
	strs := []string{"a", "b", "c", "d"}
	for _, s := range strs {
		res = append(res, func() {
			log(s)
		})
	}

	for _, f := range res {
		f()
	}
}

func log(s string) {
	fmt.Println("s =", s)
}
