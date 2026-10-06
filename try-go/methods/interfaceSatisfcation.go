package main

import "fmt"

type Counter struct {
	N int
}

func (c *Counter) Inc() {
	c.N++
}

type Incrementor interface {
	Inc()
}

func bump(i Incrementor) {
	i.Inc()
}

func main() {
	c := Counter{
		N: 0,
	}

	//bump(c); -> compile error

	bump(&c)
	fmt.Println(c)
}

/*
T implementing an interaface with a pointer method means
T does not satisfy it but *T can,

*/
