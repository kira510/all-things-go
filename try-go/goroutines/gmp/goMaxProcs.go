package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("Processors:", runtime.GOMAXPROCS(0))
	fmt.Println("CPU:", runtime.NumCPU())
}

/*
Go routines: 1
Processors: 8
CPU: 8
*/
