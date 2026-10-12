package main

import "fmt"

func main() {
	var pStock *int
	var pPrice *float32
	var pItem *struct{ Code string }

	fmt.Println(pStock, pPrice, pItem)

	fmt.Println(*pStock) //runtime error
}

/*
<nil> <nil> <nil>
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x104a22e60]

goroutine 1 [running]:
main.main()
        /Users/kiran/Documents/personal/all-things-go/try-go/pointers/basics/nilPointer.go:12 +0x70
exit status 2
*/
