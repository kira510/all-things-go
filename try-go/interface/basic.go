package main

import (
	"errors"
	"fmt"
)

type Card struct {
	Number  int
	Balance float32
}

func (c *Card) Pay(amount float32) error {
	if c.Balance < amount {
		return errors.New("Insufficient Balance")
	}

	c.Balance -= amount

	fmt.Println("updated %s card blance $.2%f", c.Number, c.Balance)

	return nil
}

type Wallet struct {
	Email   string
	Balance float32
}

func (w *Wallet) Pay(amount float32) error {
	if w.Balance < amount {
		return errors.New("Insufficient Balance in wallet")
	}

	w.Balance -= amount

	fmt.Println("updated %s wallet blance $.2%f", w.Email, w.Balance)

	return nil
}

type Payable interface {
	Pay(amount float32) error
}

func makePayment(p Payable, amount float32) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	card := Card{
		Number:  1213423,
		Balance: 3400.00,
	}

	wallet := Wallet{
		Email:   "someemail.com",
		Balance: 45000.00,
	}

	makePayment(&card, 300)
	makePayment(&wallet, 4500)
}

/*
Output
updated %s card blance $.2%f 1213423 3100
updated %s wallet blance $.2%f someemail.com 40500

------------------------
Go uses implicit (structural) interface satisfaction:
any type whose method set has Pay(float64) error is a Payable, with no implements keyword.

*/

/*
Advantages:

Polymorphism:
one piece of code works with many types at runtime.

Decoupling:
code that depends on an interface doesn't depend on any particular concrete type,
so the types can change without touching the dependent code.

*/
