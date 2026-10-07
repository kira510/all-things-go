package main

import "fmt"

type Order struct {
	Name  string
	Price float32
}

func (o Order) ApplyDiscount(percent float32) {
	o.Price = o.Price - o.Price*percent
}

type Discounter interface {
	ApplyDiscount(percent float32)
}

func main() {
	o := Order{
		Name:  "water bottle",
		Price: 100.00,
	}

	var d Discounter = o
	d.ApplyDiscount(0.1)

	fmt.Println(o)
	fmt.Println(&o)
}

/*
Output:
{water bottle 90}
&{water bottle 90}
*/
