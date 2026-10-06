/*
demonstrate passing pointer and by value
*/
package main

import "fmt"

type Product struct {
	Name string
}

// RenameByPassByVal has a value receiver: p is a copy of the struct,
// so the change only affects the copy and the original is unchanged.
func (p Product) RenameByPassByVal(name string) {
	p.Name = name
}

// RenameByRef has a pointer receiver: p points to the original struct,
// so the change persists.
// Rule of thumb: use a pointer receiver when the method mutates the
// receiver or the struct is large; use a value receiver for small,
// read-only methods.
func (p *Product) RenameByRef(name string) {
	p.Name = name
}

func main() {
	p := Product{
		Name: "Book",
	}

	p.RenameByPassByVal("Ball")

	fmt.Println(p.Name)

	// p is addressable, so Go rewrites this to (&p).RenameByRef("Bat")
	p.RenameByRef("Bat")
	fmt.Println(p.Name)
}

/*
Output:
Book
Bat
*/
