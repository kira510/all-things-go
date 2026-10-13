package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchProduct(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("fetched product:", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fetchProduct(i, &wg)
	}

	wg.Wait()
}

/*
fetched product: 1
fetched product: 2
fetched product: 0
*/
