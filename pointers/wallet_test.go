package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := Bitcoin(10)

		assetBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		// not 100% clear on this
		asserNoError(t, err)
		assetBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrorInsufficientFunds)
		assetBalance(t, wallet, Bitcoin(20))
	})
}

func assetBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	fmt.Printf("wallet mem address: %p \n", &wallet)

	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func asserNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error, but didnt want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error, but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

/*
* Accoring to lesson plan, nil is synonymous with Null, I have however heard otherwise
* Like null, if you try to access a value that is nil, it ill throw a runtime panic.
 */
