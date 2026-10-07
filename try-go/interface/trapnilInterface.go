/*
The trap is the case where the interface variable is not nil, but the concrete value inside it is nil.
*/

package main

import "fmt"

type Card struct {
	Number  int
	Balance float32
}

func (c *Card) Pay(amount float32) {
	c.Balance -= amount
}

type Payable interface {
	Pay(amount float32)
}

func main() {
	var c *Card

	var p Payable = c

	fmt.Println(c)
	fmt.Println(c == nil)
	fmt.Println(p == nil)
	fmt.Printf("%T\n", p)
}

/*
Output:

<nil>
true
false
*main.Card

*/
