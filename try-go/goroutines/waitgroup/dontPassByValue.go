package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchProduct(i int, wg sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("fetched product:", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fetchProduct(i, wg)
	}

	wg.Wait()
}

/*
fetched product: 8
fetched product: 7
fetched product: 9
fetched product: 6
fetched product: 5
fetched product: 3
fetched product: 2
fetched product: 4
fetched product: 1
fetched product: 0
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x140000021a0?)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/runtime/sema.go:62 +0x2c
sync.(*WaitGroup).Wait(0x1400011e020)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/sync/waitgroup.go:116 +0x74
main.main()
        /Users/kiran/Documents/personal/all-things-go/try-go/goroutines/waitgroup/dontPassByValue.go:23 +0x9c
exit status 2
*/

/*
Each go fetchProduct(i, wg) hands the goroutine a copy of the WaitGroup.
wg.Done() decrements the copy’s counter, and the original in main stays at 10.
After all goroutines finish, wg.Wait() is still waiting for a count that never drops,
and no goroutine is left to wake it, so Go reports a deadlock.

*/

/*
Go vet will flag this:


goroutines git:(main) ✗ go vet waitgroup/dontPassByValue.go
# command-line-arguments
waitgroup/dontPassByValue.go:9:29: fetchProduct passes lock by value: sync.WaitGroup contains sync.noCopy
waitgroup/dontPassByValue.go:20:22: call of fetchProduct copies lock value: sync.WaitGroup contains sync.noCopy

*/
