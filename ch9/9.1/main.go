package main

import (
	"fmt"
	"log"

	"go-book/ch9/9.1/bank"
)

func main() {
	b, err := bank.NewBank("ICPC bank")
	if err != nil {
		log.Fatalln(err)
	}

	a1, err := b.NewAccount("xiaoming")
	if err != nil {
		log.Fatalln(err)
	}
	go a1.Monitor()

	a2, err := b.NewAccount("xiaohong")
	if err != nil {
		log.Fatalln(err)
	}
	go a2.Monitor()

	fmt.Println("number of bank =", b.NumberOfAccount())

	go func() {
		a1.Deposit(10)
		fmt.Println("after Deposit, balance of a1 =", a1.Balance())
	}()

	go func() {
		a1.Withdraw(10)
		fmt.Println("after withdraw, balance of a1 =", a1.Balance())
	}()

	select {}
}
