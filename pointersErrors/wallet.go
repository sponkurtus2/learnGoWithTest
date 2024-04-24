package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC ", b)
}

type Wallet struct {
	money Bitcoin
}

func (w *Wallet) Deposit(moneyToDeposit Bitcoin) {
	// return newAmount
	w.money += moneyToDeposit
}

var ErrInsufficientFunds = errors.New("cant withdraw,insufficient funds")

func (w *Wallet) Whithdraw(moneyToWithDraw Bitcoin) error {
	if moneyToWithDraw > w.Balance() {
		return ErrInsufficientFunds
	}

	w.money -= moneyToWithDraw
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.money
}
