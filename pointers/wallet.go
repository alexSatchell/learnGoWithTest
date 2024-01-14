package pointers

import (
	"errors"
	"fmt"
)

// Go allows for creating of new types (im guessing alias's)
// from existing ones
type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go, when you call a function (like most langs ik of)
// the argunmemtns are copied

type Wallet struct {
	// if a symbol is lowercased like balance is,
	// it will be private, only accessible by methods
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address if balance in deposit is %p \n", &w.balance)
	// In go, we do not need to explicitly de-referance
	// furthermore, struct pointers are auto de-referenced
	w.balance += amount
}

// var keyword will be a global variable to the package
var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
