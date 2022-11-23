package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(&a)
	fmt.Println(a)
}

func reverse(slice *[]int) {
	ptr := *slice
	n := len(ptr)
	for i := 0; i < n/2; i++ {
		ptr[i], ptr[n-1-i] = ptr[n-1-i], ptr[i]
	}
}
