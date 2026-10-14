package main

import "fmt"

func returnProcessedOrder(orderCh chan int) {
	orderCh <- 100
}

func main() {
	orders := make(chan int)

	go returnProcessedOrder(orders)

	i := <-orders
	fmt.Println("Processed order", i)
}
