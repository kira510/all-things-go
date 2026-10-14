package main

import (
	"errors"
	"fmt"
	"sync"
)

type Cart struct {
	mu    sync.Mutex
	items map[string]int
}

func NewCart() *Cart {
	return &Cart{items: make(map[string]int)}
}

func (c *Cart) AddItem(item string, quantity int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if quantity <= 0 {
		return errors.New("quanity cant be zero")
	}

	c.items[item] += quantity

	return nil
}

func main() {
	cart := NewCart()

	if err := cart.AddItem("pen", 5); err != nil {
		fmt.Println("add failed:", err)
	}

	if err := cart.AddItem("Book", -1); err != nil {
		fmt.Println("add failed:", err)
	}

	fmt.Println("Cart:", cart.items)
}
