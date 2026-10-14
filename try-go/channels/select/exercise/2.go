package main

import "fmt"

func logAdded(item string) {
	fmt.Println("added", item)
}

func logDropped(item string) {
	fmt.Println("cart full, dropping", item)
}

func main() {
	cart := make(chan string, 2)
	items := []string{"shoes", "hat", "scarf"}

	for _, item := range items {
		select {
		case cart <- item:
			logAdded(item)
		default:
			logDropped(item)
		}
	}
}
