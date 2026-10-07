package main

import "fmt"

type Payable interface {
	Pay()
}

func main() {
	var p Payable

	fmt.Println(p == nil)
}

/*
Output:

true
*/
