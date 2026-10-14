package main

import (
	"fmt"
	"time"
)

func requestQuote(delay time.Duration) chan string {
	quote := make(chan string)

	go func() {
		defer close(quote)
		time.Sleep(delay)
		quote <- "24.99"
	}()

	return quote
}

func printQuote(quote string) {
	fmt.Println("quote:", quote)
}

func printTimeout(timeout time.Duration) {
	fmt.Println("timeout after", timeout)
}

func main() {
	quoteCh := requestQuote(500 * time.Millisecond)
	timeOut := 10 * time.Millisecond

	select {
	case quote := <-quoteCh:
		printQuote(quote)
	case <-time.After(timeOut):
		printTimeout(timeOut)

	}
}
