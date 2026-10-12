package main

import "fmt"

func main() {
	done := make(chan struct{})
	status := "queued"

	go func(statusParam string) {
		fmt.Println(statusParam)
		fmt.Println(status)
		close(done)
	}(status)

	status = "no queue"
	<-done
}

/*
queued
no queue
*/
