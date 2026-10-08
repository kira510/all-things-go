package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	raw := `{"product": "BOOK-01", "stock": 12, "in_promotion": true}`

	var decode map[string]interface{}

	if err := json.Unmarshal([]byte(raw), &decode); err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(decode)
	fmt.Println(&decode)
}

/*

map[in_promotion:true product:BOOK-01 stock:12]
&map[in_promotion:true product:BOOK-01 stock:12]

*/
