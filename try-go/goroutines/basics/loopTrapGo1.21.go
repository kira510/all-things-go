package main

import (
	"fmt"
	"time"
)

func main() {
	orders := []int{123, 45, 453}

	for i := 0; i < len(orders); i++ {
		go func() {
			fmt.Println("first loop:", i)
		}()
	}

	time.Sleep(1000 * time.Microsecond)

	for i := 0; i < len(orders); i++ {
		go func(k int) {
			fmt.Println("second loop:", k)
		}(i)
	}

	time.Sleep(1000 * time.Microsecond)
}

/*

first loop: 3
first loop: 3
first loop: 3
second loop: 0
second loop: 1
second loop: 2

*/

/*
// On Go 1.22+, this works correctly.
	// On Go 1.21 and earlier, all three goroutines see i == 3.
*/
