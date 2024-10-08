package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("cant with draw")

// Create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}

	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount

	return nil
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}
