package main

import "fmt"

func main() {
	smallCh := make(chan int, 2)
	largeCh := make(chan int, 3)

	smallCh <- 3
	smallCh <- 4
	close(smallCh)

	largeCh <- 5
	largeCh <- 5
	largeCh <- 10
	close(largeCh)

	smallCount := 0
	largeCount := 0
	active := 2

	for active > 0 {
		select {
		case val, ok := <-smallCh:
			if !ok {
				smallCh = nil
				active--
				continue
			}
			smallCount += val
		case val, ok := <-largeCh:
			if !ok {
				largeCh = nil
				active--
				continue
			}
			largeCount += val
		}
	}

	fmt.Println("small sum:", smallCount)
	fmt.Println("large sum:", largeCount)
	fmt.Println("grand total:", smallCount+largeCount)

}
