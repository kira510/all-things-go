package main

import "errors"

type Card struct {
	Number  int
	Balance float32
}

func (c *Card) Pay(amount float32) error {
	if c.Balance < amount {
		return errors.New("insuffient balance")
	}

	c.Balance -= amount

	return nil
}

type Payable interface {
	Pay(amount float32) error
}

// Compile-time assertion: *CreditCard must satisfy Payable.
var _ Payable = (*Card)(nil)
var _ Payable = &Card{} //for *T, if T satisfies then user: var _ Payable = Card{}

func main() {
	c := Card{
		Number:  2142343,
		Balance: 6788,
	}

	c.Pay(455)
}

/*
Extremely used pattern.

The standard library uses this pattern heavily.
Search for var _ io.Writer = in the Go source tree and there are dozens of examples.

*/
