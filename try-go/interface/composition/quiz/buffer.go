package main

import (
	"bytes"
	"fmt"
	"io"
)

func echo(rw io.ReadWriter) {
	_, _ = rw.Write([]byte("go"))
	p := make([]byte, 2)
	n, _ := rw.Read(p)

	fmt.Printf("%s %d\n", p[:n], n)
	fmt.Printf("%s\n", p)
}

func main() {
	var b bytes.Buffer
	echo(&b)
}

/*
go 2
*/
