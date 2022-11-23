package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum =", Sum(nums...))
	max, _ := Max(nums...)
	fmt.Println("max =", max)
	min, _ := Min(nums...)
	fmt.Println("min =", min)
}

func Sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}

	return sum
}

func Max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty numbers")
	}
	max := math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max, nil
}

func Min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty numbers")
	}
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min, nil
}
