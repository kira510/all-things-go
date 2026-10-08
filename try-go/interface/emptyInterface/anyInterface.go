package main

import "fmt"

func main() {
	var a any = 1
	var b interface{} = "string val"

	fmt.Println(a, b)
}

/*
Go 1.18 introduced any as a predeclared alias for interface{}.
They are exactly the same type.


The old name still works, and is interface{} constantly in older code
and in the standard library, but
		gofmt -r 'interface{} -> any'
is a one-liner that converts a codebase, and most teams have made the switch.

*/
