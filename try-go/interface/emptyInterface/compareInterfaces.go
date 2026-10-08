package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	var a interface{} = 42
	var b interface{} = 42
	var c interface{} = int64(42)

	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)

	var n interface{} = []int{1, 2, 3}
	var m interface{} = []int{1, 2, 3}

	fmt.Println("n == m", n == m)
}

/*
a == b true
a == c false
Recovered:  runtime error: comparing uncomparable type []int
*/

/*
Two any values are equal when both their dynamic types and dynamic values are equal.


Slices, maps, and functions aren't comparable with == in Go (except against nil).
When they're wrapped in any, the comparison compiles, but the runtime check fails.

*/
