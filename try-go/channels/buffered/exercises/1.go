package main

import "fmt"

func sendOrders(orderCh chan string) {
	defer close(orderCh)
	orderCh <- "order-101"
	fmt.Println("sent 101")

	orderCh <- "order-102"
	fmt.Println("sent 102")

	orderCh <- "order-103"
	fmt.Println("sent 103")

	orderCh <- "order-104"
	fmt.Println("sent 104")

	fmt.Printf("order ch queued=%d capacity=%d\n", len(orderCh), cap(orderCh))
}

func main() {
	orders := make(chan string, 3)

	go sendOrders(orders)

	fmt.Printf("queued=%d capacity=%d\n", len(orders), cap(orders))

	for order := range orders {
		fmt.Println("ship:", order)
	}
}

/*
go run exercises/1.go
queued=0 capacity=3
ship: order-101
sent 101
sent 102
sent 103
sent 104
order ch queued=2 capacity=3
ship: order-102
ship: order-103
ship: order-104
*/
