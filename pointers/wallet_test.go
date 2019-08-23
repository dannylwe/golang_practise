package wallet

import (
	"errors"
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("test deposit on wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		if got != want {
			t.Errorf("got %v but was expecting balance of %v", got, want)
		}
	})
	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 200}
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := 190


		if got != want {
			t.Errorf("got %v but was expecting balance of %v", got, want)
		}
	})
	t.Run("test account has sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(1000)
		
		var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
		if err == nil {
			t.Errorf("got %v but wanted %v", err, ErrInsufficientFunds)
		}
	})

}
