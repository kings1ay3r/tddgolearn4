package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
	fmt.Printf("address of balance in test is %v \n", &w.balance)
}

func (w *Wallet) Balance() int {
	return w.balance
}
