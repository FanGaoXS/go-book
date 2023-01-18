package bank

import "sync"

var (
	balance = 0
	lock    sync.Mutex
)

func Deposit(amount int) {
	lock.Lock()
	defer lock.Unlock()
	balance += amount
}

func Balance() int {
	lock.Lock()
	defer lock.Unlock()
	return balance
}
