package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient-balance")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	// fmt.Printf("address of balance in test is %v \n", &w.balance)
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
	// fmt.Printf("address of balance in test is %v \n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
