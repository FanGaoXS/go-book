package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(naturals, squarers)

	for x := range squarers {
		fmt.Println(x)
	}
}

func counter(out chan<- int) {
	i := 0
	for i <= 5 {
		out <- i
		i++
		time.Sleep(time.Second * 1)
	}
	close(out)
}

func squarer(in chan int, out chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
