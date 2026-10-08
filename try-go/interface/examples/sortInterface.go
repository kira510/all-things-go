package main

import (
	"fmt"
	"sort"
)

type ByAge []int

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i] < b[j]
}

func main() {
	ages := ByAge{
		10, 2, 39, 34, 27,
	}

	sort.Sort(ages)
	fmt.Println(ages)
}
