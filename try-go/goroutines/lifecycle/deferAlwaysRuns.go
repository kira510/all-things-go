package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)                     //defer last
		defer fmt.Println("cleanup")          // defer second
		defer fmt.Println("starting cleanup") //defer first
		fmt.Println("start")
		runtime.Goexit()
		fmt.Println("after")
	}()
	<-done
}

/*
start
cleanup
*/

/*
Close will not return and end main,
it will return but before that all defer will be executed

Defer is a stack, so last in first out. LIFO
*/
