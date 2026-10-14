package main

import "fmt"

func receiveReport(label string, numbers <-chan int) string {
	value, ok := <-numbers

	return fmt.Sprintf("%s: value=%d ok=%t len=%d cap=%d", label, value, ok, len(numbers), cap(numbers))
}

func main() {
	numbers := make(chan int, 2)
	numbers <- 11
	numbers <- 22
	close(numbers)

	fmt.Println(receiveReport("receive 1", numbers))
	fmt.Println(receiveReport("receive 2", numbers))
	fmt.Println(receiveReport("receive 3", numbers))
}
