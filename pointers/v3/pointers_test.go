package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, b Bitcoin) {

		got := w.balance
		want := b
		if got != b {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(10)

		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(15)
		wallet.Withdraw(5)
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
}
