package main

import "fmt"

type Event struct {
	EventName string
	Payload   interface{} //can hold multiple values
}

func (e *Event) send() {
	fmt.Printf("Event %s : %v\n", e.EventName, e.Payload)
}

func main() {
	events := []Event{
		{EventName: "product_viewed", Payload: "Book"},
		{EventName: "order_placed", Payload: 24.00},
	}

	for _, event := range events {
		event.send()
	}
}

/*
Output

Event product_viewed : Book
Event order_placed : 24
*/
