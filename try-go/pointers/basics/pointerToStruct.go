package main

import "fmt"

type Product struct {
	Code  string
	Stock int
}

func main() {
	book := Product{Code: "BOOK-01", Stock: 12}
	pStock := &book.Stock
	*pStock = 9
	other := book
	other.Stock = 1
	fmt.Println(book.Stock, other.Stock)
}

/*
9 1
*/
