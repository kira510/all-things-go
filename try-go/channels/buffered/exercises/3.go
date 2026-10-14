package main

import "fmt"

func tryEnqueue(queue chan<- string, message string) bool {
	select {
	case queue <- message:
		return true
	default:
		return false
	}
}

func main() {
	queue := make(chan string, 2)
	messages := []string{
		"pen", "book", "ball",
	}

	for _, msg := range messages {
		if tryEnqueue(queue, msg) {
			fmt.Println("enqueued:", msg)
		} else {
			fmt.Println("dropped:", msg)
		}
	}

	close(queue)
	for item := range queue {
		fmt.Println(item)
	}
}

/*
enqueued: pen
enqueued: book
dropped: ball
pen
book
*/
