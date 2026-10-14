package main

import "fmt"

func printHandled(msg string) {
	fmt.Println("handled:", msg)
}

func main() {
	orders := make(chan string, 1)
	refunds := make(chan string, 1)
	refunds <- "refund #9 approved"

	select {
	case msg := <-orders:
		printHandled(msg)
	case msg := <-refunds:
		printHandled(msg)
	}
}
