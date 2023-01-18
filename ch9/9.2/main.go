package main

import (
	"fmt"
	"go-book/ch9/9.2/bank"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			do()
		}()
	}
	for {
	}
}

func do() {
	fmt.Println("before deposit, balance is", bank.Balance())
	bank.Deposit(10)
	fmt.Println("after deposit, deposit is", bank.Balance())
}
