package main

import (
	"fmt"
	"runtime"
	"time"
)

func ProcessOrder(i int) {
	fmt.Println("Processing:", i)
}

func main() {
	fmt.Println("Goroutines at start:", runtime.NumGoroutine())

	for i := 0; i < 5; i++ {
		go ProcessOrder(i)
	}

	fmt.Println("Goroutines at initilization:", runtime.NumGoroutine())

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Goroutines at end:", runtime.NumGoroutine())
}

/*
Goroutines at start: 1
Goroutines at initilization: 6
Processing: 4
Processing: 3
Processing: 0
Processing: 2
Processing: 1
Goroutines at end: 1
*/
