package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcessItem(orderId int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Item Processed:", orderId)
}

func main() {
	orderBatch := [][]int{
		{123, 456, 567},
		{44, 55, 55},
	}

	var wg sync.WaitGroup
	for batchNo, batch := range orderBatch {
		fmt.Println("Processing batch:", batchNo)
		for _, item := range batch {
			// wg.Go(ProcessItem(item)), Calls ADD, then defer go too, no need to call defer wg.done

			go ProcessItem(item, &wg)
		}

		wg.Wait()
	}
}

/*

Processing batch: 0
Item Processed: 456
Item Processed: 567
Item Processed: 123
Processing batch: 1
Item Processed: 55
Item Processed: 44
Item Processed: 55

*/

/*
Requires Go 1.25+.

wg.Go(func) does three things internally: it calls Add(1),
it spawns a goroutine, and it arranges for Done to run when the function returns (whether normally or via panic).
The boilerplate disappears, and so do the bugs around forgetting Add or Done.

*/
