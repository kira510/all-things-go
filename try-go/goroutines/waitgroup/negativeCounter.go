package main

import (
	"fmt"
	"sync"
	"time"
)

func processItem(id int, wg *sync.WaitGroup) {
	wg.Done()
	wg.Done() // Bug: double Done
	time.Sleep(1000 * time.Millisecond)

	fmt.Println(id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go processItem(i, &wg)
	}

	wg.Wait()
}

/*

go run waitgroup/negativeCounter.go
panic: sync: negative WaitGroup counter

goroutine 34 [running]:
sync.(*WaitGroup).Add(0x0?, 0x0?)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/sync/waitgroup.go:62 +0x108
sync.(*WaitGroup).Done(...)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/sync/waitgroup.go:87
main.processItem(0x0?, 0x0?)
        /Users/kiran/Documents/personal/all-things-go/try-go/goroutines/waitgroup/negativeCounter.go:10 +0x30
created by main.main in goroutine 1
        /Users/kiran/Documents/personal/all-things-go/try-go/goroutines/waitgroup/negativeCounter.go:22 +0x38
exit status 2

*/

/*

It’s a race, so it may or may not panic.

There are 2 Adds but 4 Dones, so the counter can go negative, which gives panic: sync: negative WaitGroup counter.
But goroutine 0 can run both its Dones before goroutine 1 starts.
The counter goes 2, 1, 0, so it never goes negative, and wg.Wait() is released.
main then returns and the program exits, so goroutine 1 may never run its Dones.
The time.Sleep and Println are never reached either way, so nothing prints.

*/
