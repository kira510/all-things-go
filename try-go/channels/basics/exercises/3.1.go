package main

import (
	"fmt"
	"time"
)

func returnProcessedOrder(orderCh chan int) {
	time.Sleep(5000 * time.Millisecond)
	for i := 0; i <= 1; i++ {
		orderCh <- i
	}
}

func main() {
	orders := make(chan int)

	go returnProcessedOrder(orders)

	i := <-orders
	fmt.Println("Processed order", i)
}
