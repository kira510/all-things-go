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
			wg.Add(1)
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
