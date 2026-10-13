package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			wg.Add(1)
			time.Sleep(10 * time.Millisecond) // keep the counter non-zero while main wakes
		}()
	}

	wg.Wait() // wakes up, sees the counter is no longer 0, panics
	fmt.Println("ok")
}

/*

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x140000021a0?)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/runtime/sema.go:62 +0x2c
sync.(*WaitGroup).Wait(0x14000110020)
        /opt/homebrew/Cellar/go@1.21/1.21.7/libexec/src/sync/waitgroup.go:116 +0x74
main.main()
        /Users/kiran/Documents/personal/all-things-go/try-go/goroutines/waitgroup/addAfterWaitPanic.go:22 +0x98
exit status 2
*/
