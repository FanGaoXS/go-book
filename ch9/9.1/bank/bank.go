package bank

import (
	"errors"

	"github.com/google/uuid"
)

var (
	banks = map[string]*Bank{} // name -> bank
)

type Bank struct {
	id       string // uuid
	name     string
	accounts map[string]*Account
}

func NewBank(name string) (*Bank, error) {
	if _, ok := banks[name]; ok {
		return nil, errors.New("bank 已经存在")
	}

	b := &Bank{
		id:       uuid.New().String(),
		name:     name,
		accounts: make(map[string]*Account),
	}
	banks[name] = b
	return b, nil
}

func (b Bank) NumberOfAccount() int {
	return len(b.accounts)
}

func (b *Bank) NewAccount(name string) (*Account, error) {
	if _, ok := b.accounts[name]; ok {
		return nil, errors.New("name 存在")
	}

	account, err := newAccount(name, b.name)
	if err != nil {
		return nil, err
	}
	b.accounts[name] = account

	return account, nil
}

type Amount struct {
	isDeposit bool
	number    int
}

type Account struct {
	name string

	bankId    string
	amountCh  chan *Amount
	balanceCh chan int
}

func newAccount(name string, bankName string) (*Account, error) {
	if _, ok := banks[bankName]; !ok {
		return nil, errors.New("bank 不存在")
	}

	amountCh := make(chan *Amount)
	balanceCh := make(chan int)
	return &Account{
		name:      name,
		bankId:    banks[bankName].id,
		amountCh:  amountCh,
		balanceCh: balanceCh,
	}, nil
}

func (a *Account) Monitor() {
	balance := 0
	for {
		select {
		case amount := <-a.amountCh:
			if amount.isDeposit {
				balance += amount.number
				continue
			}
			balance -= amount.number
		case a.balanceCh <- balance:
		}
	}
}

func (a *Account) Deposit(amount int) {
	am := &Amount{
		isDeposit: true,
		number:    amount,
	}

	a.amountCh <- am
}

func (a *Account) Withdraw(amount int) {
	am := &Amount{
		isDeposit: false,
		number:    amount,
	}

	a.amountCh <- am
}

func (a *Account) Balance() int {
	return <-a.balanceCh
}
