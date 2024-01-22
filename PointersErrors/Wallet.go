package PointersErrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", b)

}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {

	w.balance += deposit

}

func (w *Wallet) Balance() Bitcoin {

	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if w.balance < amount {

		return errors.New("cant withdraw that much")
	}
	w.balance -= amount
	return nil
}
