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
	prices := map[string]Cart{
		"notebook": Cart{Name: "book", Quantity: 1},
	}

	//prices["notebook"].AddOne()  //cannot call pointer method AddOne on Cart
	fmt.Println(prices)

	//How to call?
	prices2 := map[string]*Cart{
		"novel": &Cart{
			Name: "Lord of rigs", Quantity: 1,
		},
	}
	prices2["novel"].AddOne() //compilies

	//How to mutate first one? Copy, mutate and paste back
	item := prices["notebook"]
	item.AddOne()
	prices["notebook"] = item
	fmt.Println(prices["notebook"])
}

/*
Expression	                                     Addressable?
Named local or package variable (item)	         Yes
Field of an addressable struct (order.Item)	     Yes
Element of a slice (items[3])	                   Yes
Pointer dereference (*p)	                       Yes

Element of a map (prices["mouse"])	             No
Return value from a function (getItem())	       No
Result of a type conversion (Product(other))	   No
Composite literal value (CartItem{...})	         No
*/
