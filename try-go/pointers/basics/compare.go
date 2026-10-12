package main

import "fmt"

func main() {
	a := 10
	b := 10

	pA := &a
	pB := &b
	pA2 := &a

	fmt.Println("pA == pB: ", pA == pB)
	fmt.Println("pA == pA2:", pA == pA2)
	fmt.Println("*pA == *pB:", *pA == *pB)
}

/*
pA == pB:  false
pA == pA2: true
*pA == *pB: true
*/
