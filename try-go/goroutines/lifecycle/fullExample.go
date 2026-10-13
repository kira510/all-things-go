package main

import (
	"fmt"
	"runtime"
	"time"
)

func processOrder(id int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Order %d: recovered from %v\n", id, r)
		}
	}()

	fmt.Printf("Order %d: starting\n", id)
	time.Sleep(100 * time.Millisecond) // Running -> Blocked -> Runnable
	fmt.Printf("Order %d: awake\n", id)

	if id == 2 {
		panic("bad order data")
	}

	fmt.Printf("Order %d: done\n", id)
}

func main() {
	fmt.Println("Before launch:", runtime.NumGoroutine(), "goroutines")

	for i := 1; i <= 3; i++ {
		go processOrder(i)
	}

	fmt.Println("After launch: ", runtime.NumGoroutine(), "goroutines")

	time.Sleep(300 * time.Millisecond) // wait for them all to finish

	fmt.Println("After wait:  ", runtime.NumGoroutine(), "goroutines")
}

/*
Before launch: 1 goroutines
After launch:  4 goroutines
Order 3: starting
Order 1: starting
Order 2: starting
Order 2: awake
Order 1: awake
Order 1: done
Order 3: awake
Order 3: done
Order 2: recovered from bad order data
After wait:   1 goroutines
*/
