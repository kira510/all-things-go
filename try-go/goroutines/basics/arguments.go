package main

import (
	"fmt"
	"time"
)

func showProduct(name string, price float32) {
	fmt.Printf("Routine print: Product: %s has price: $%f\n", name, price)
}

func main() {
	product := "book"
	var price float32 = 24.99

	go showProduct(product, price)

	product = "pen"
	price = 15.99

	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Product: %s has price: $%f\n", product, price)
}

/*
Routine print: Product: book has price: $24.990000
Product: pen has price: $15.990000
*/

/*
The goroutine got copies of product and price at the moment the go call ran,
so the later reassignments to those variables have no effect on what the goroutine prints.
*/
