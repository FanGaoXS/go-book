package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	res := rotate(a, 9)
	fmt.Println(res)
}

func rotate(s []int, position int) []int {
	if position <= 0 || position > len(s) {
		return nil
	}

	index := position - 1
	res := s[index+1:]
	res = append(res, s[index])

	for i := 0; i < index; i++ {
		res = append(res, s[i])
	}
	return res
}
