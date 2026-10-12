package main

import "fmt"

type Book struct {
	Code  string
	Price float32
	Stock int
}

func main() {
	a := 5
	k := &a

	fmt.Println(a)
	fmt.Println(k)

	b := Book{
		Code:  "123",
		Price: 12.99,
		Stock: 100,
	}

	pBook := &b
	pPrice := &b.Price // pointer to one field
	pStock := &b.Stock //pointer to another field

	fmt.Printf("&book:       %p\n", pBook)
	fmt.Printf("&book.Price: %p\n", pPrice)
	fmt.Printf("&book.Stock: %p\n", pStock)

	fmt.Printf("&book:       %v\n", pBook)
	fmt.Println(pBook)
}

/*
5
0x14000112008
&book:       0x14000126000
&book.Price: 0x14000126010
&book.Stock: 0x14000126018
&book:       &{123 12.99 100}
&{123 12.99 100}
*/
