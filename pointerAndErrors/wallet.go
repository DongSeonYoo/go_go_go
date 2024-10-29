package pointer

import (
	"errors"
	"fmt"
)

/*
	Pointer.
	- Go copies values when you pass them functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
	- The facts that Go takes a copy of values is useful a lot of the time, but somethimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

	Nil
	- Pointers can be nil
	- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runime exception - the compiler won't help you here.
	- Useful for when you want to describe a value that could be missing

	Errors
	- Errors are the way to signify failure whe calling a function/method
	- By listening to our tests we concluded that checking for a string in an error would result in a flaky test
	- *This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
*/

var ErrInsufficientFunds = errors.New("You don't have enough balance")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

// When call a function or a method the arguments are copied.
// When calling Deposit. the w is copy of whatever we called the method from.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
