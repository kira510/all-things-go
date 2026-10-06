/*
demonstrate how arrays and slices inside a struct behave with
value receivers vs pointer receivers
*/
package main

import "fmt"

// Note: the field names are swapped compared to their types.
// scoreSlice is actually an ARRAY ([3]int): the numbers sit inside the struct.
// scoreArray is actually a SLICE ([]int): only a header (pointer, len, cap)
// sits inside the struct, and the numbers live in a separate backing array.
type Scores struct {
	scoreSlice [3]int
	scoreArray []int
}

// BumpFirst has a value receiver: s is a copy of the whole struct.
//   - scoreSlice (array): the copy has its own 3 numbers, so only the copy
//     changes and the caller does not see it.
//   - scoreArray (slice): the copy has the same pointer to the backing array,
//     so the change goes to the shared data and the caller sees it.
func (s Scores) BumpFirst() {
	s.scoreSlice[0]++
	s.scoreArray[0]++
}

// BumpFirstWithPointer has a pointer receiver: s points to the original
// struct, so both changes are made to the caller's data.
func (s *Scores) BumpFirstWithPointer() {
	s.scoreSlice[0]++
	s.scoreArray[0]++
}

func main() {
	s := Scores{
		scoreSlice: [3]int{10, 20, 30},
		scoreArray: []int{40, 50, 60},
	}

	// array stays 10 (copy changed), slice goes 40 -> 41 (shared data changed)
	s.BumpFirst()
	fmt.Println(s)

	// both change: array 10 -> 11, slice 41 -> 42
	s.BumpFirstWithPointer()
	fmt.Println(s)
}

/*
Output:

{[10 20 30] [41 50 60]}
{[11 20 30] [42 50 60]}

*/
