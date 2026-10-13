package main

import (
	"fmt"
	"time"
)

func ProcessOrder(orderId string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if orderId == "" {
		panic("empty order id \n")
	}

	fmt.Printf("processed: %s \n", orderId)
}

func main() {
	go ProcessOrder("123")
	go ProcessOrder("")
	go ProcessOrder("345")

	time.Sleep(1000 * time.Millisecond)
}

/*

processed: 345
processed: 123
empty order id

*/

/*
defer recover() Inside the Goroutine
If you want a single bad input to fail one goroutine without killing the rest,
wrap the work in a defer recover()
*/
