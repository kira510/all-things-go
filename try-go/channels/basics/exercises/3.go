package main

import (
	"fmt"
	"time"
)

func processOrders(orderCh chan int) {
	defer close(orderCh)
	for i := 0; i <= 1; i++ {
		time.Sleep(1000 * time.Millisecond)
		orderCh <- i
	}
}

func main() {
	orderCh := make(chan int, 2)

	go processOrders(orderCh)

	for k := range orderCh {
		fmt.Println(k)
	}
}
