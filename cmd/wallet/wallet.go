package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amnt Bitcoin) {
	w.balance += amnt
}

func (w *Wallet) Withdrawl(amnt Bitcoin) error {
	if amnt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amnt
	return nil
}

// passed in pointer because other wise it would have passed in the value or copy of the original wallet
// by passing pointer we have passed the memory address or reference to the original wallet object.
func (w *Wallet) Balance() Bitcoin {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	return w.balance
}
