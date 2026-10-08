package main

import "fmt"

func main() {
	var x interface{} = 19.999

	//cant do operations on interface without infering type -> complile error
	//y := x + 90.1 //invalid operation: mismatched types interface{}

	if price, ok := x.(float64); ok {
		fmt.Printf("%0.2f\n", price)
	} else {
		fmt.Println("not a float")
	}
}

/*
20.00


To do anything beyond storing, printing with %v,
or comparing for equality, you have to recover the concrete type.
There are two ways: a type assertion, which pulls out one specific type, and a type switch,
which checks several candidates in one place.

*/
