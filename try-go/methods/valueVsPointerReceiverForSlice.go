/*
demonstrate value vs pointer receivers when the struct holds a slice.

A slice field stores only a header: (pointer to backing array, len, cap).
  - Changing an element (c.Items[i] = x) writes through the pointer into the
    shared backing array, so it works with either receiver.
  - append returns a NEW header (new len, maybe a new array). Assigning it to
    c.Items replaces the header, so it only sticks with a pointer receiver.
*/
package main

import "fmt"

type Cart struct {
	Items []string
}

// Append has a value receiver: c is a copy of the cart, header included.
// append gives a new header, which is stored in the copy and lost on return.
func (c Cart) Append(item string) {
	c.Items = append(c.Items, item)
}

// AppendByPointer has a pointer receiver: c points to the caller's cart,
// so the new header from append is stored in the original.
func (c *Cart) AppendByPointer(item string) {
	c.Items = append(c.Items, item)
}

// Rename has a pointer receiver: it writes into the backing array of the
// caller's slice.
func (c *Cart) Rename(index int, name string) {
	c.Items[index] = name
}

// RenameByValue has a value receiver, but it still works: the copied header
// has the same pointer, so this writes into the same backing array.
// It only changes an element, not the header itself.
func (c Cart) RenameByValue(index int, name string) {
	c.Items[index] = name
}

func main() {
	// zero value: c.Items is a nil slice (no array, len 0)
	var c Cart

	// both appends are made to copies, so c.Items stays nil -> {[]}
	c.Append("book")
	c.Append("glass")

	fmt.Println(c)

	// pointer receiver: c.Items gets the new headers -> {[airpods pen]}
	c.AppendByPointer("airpods")
	c.AppendByPointer("pen")

	fmt.Println(c)

	c2 := Cart{
		Items: []string{"marker", "torch"},
	}

	// element change through the pointer receiver -> {[tablets torch]}
	c2.Rename(0, "tablets")
	fmt.Println(c2)

	// element change through a copied header that has the same array pointer
	// -> {[spinner torch]}
	c2.RenameByValue(0, "spinner")
	fmt.Println(c2)
}

/*
Output:

{[]}
{[airpods pen]}
{[tablets torch]}
{[spinner torch]}

*/
