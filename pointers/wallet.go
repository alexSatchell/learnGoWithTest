package pointers

type Wallet struct {
	// if a symbol is lowercased like balance is,
	// it will be private, only accessible by methods
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
