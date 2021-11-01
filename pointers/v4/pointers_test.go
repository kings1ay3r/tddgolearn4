package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("got %q want %q", err, want)
		}
	}
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
	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(25)
		wallet.Withdraw(10)
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(15)

		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(15)
		err := wallet.Withdraw(50)
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(15)
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, want)

		if err == nil {
			t.Errorf("Wanted an error didnt get one")
		}
	})
}
