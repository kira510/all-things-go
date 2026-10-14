package main

import (
	"fmt"
	"sync"
)

type Invetory struct {
	mu       sync.Mutex
	quantity int
}

func (i *Invetory) Increment() {
	i.mu.Lock()
	i.quantity++
	i.mu.Unlock()
}

func main() {
	pens := Invetory{
		quantity: 100,
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pens.Increment()
		}()
	}

	wg.Wait()

	fmt.Println(pens.quantity)
}

/*

1100

*/
