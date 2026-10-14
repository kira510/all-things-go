package main

import "fmt"

func main() {
	var readings chan int

	fmt.Println(readings == nil)

	readings = make(chan int)

	fmt.Println(readings == nil)
}
