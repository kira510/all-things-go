package main

import "fmt"

type Cart struct {
	Name     string
	Quantity int
}

func (c *Cart) AddOne() {
	c.Quantity++
}

func main() {
	cart := []Cart{
		{Name: "pen", Quantity: 1},
		{Name: "Book", Quantity: 1},
	}

	for _, item := range cart {
		item.AddOne() //returns a copy of cart to item hence no change
	}

	fmt.Println(cart)

	for i := range cart {
		cart[i].AddOne() //change on cart itseld as range returns an addressable
	}

	fmt.Println(cart)
}

/*
Output:

[{pen 1} {Book 1}]
[{pen 2} {Book 2}]

*/
