package main

import "fmt"

type Book struct {
	Code     string
	Price    float32
	Quantity int
}

func main() {
	book := Book{
		Code:     "123",
		Price:    12.99,
		Quantity: 100,
	}

	var code *string = &book.Code
	var price *float32 = &book.Price
	var pBook *Book = &book

	fmt.Printf("pStock is %T pointing to %s\n", code, *code)
	fmt.Printf("pPrice is %T pointing to %g\n", price, *price)
	fmt.Printf("pBook  is %T pointing to %+v\n", pBook, *pBook)
	fmt.Println(code, price, pBook)
	fmt.Println(&code, &price, &pBook)
	fmt.Println(*code, *price, *pBook)

	pBook.Quantity = 99

	fmt.Println(book)
}

/*
pStock is *string pointing to 123
pPrice is *float32 pointing to 12.99
pBook  is *main.Book pointing to {Code:123 Price:12.99 Quantity:100}
0x14000124000 0x14000124010 &{123 12.99 100}
0x14000108018 0x14000108020 0x14000108028
123 12.99 {123 12.99 100}
{123 12.99 99}
*/
