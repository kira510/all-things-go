package main

import "fmt"

func main() {
	stock := 123
	pStock := &stock

	stock--
	*pStock = *pStock - 1

	fmt.Println("stock: ", stock)
	fmt.Println("*pStock:", *pStock)
}

/*
stock:  121
*pStock: 121
*/
