package main

import "fmt"

func main() {
	config := map[string]interface{}{
		"timeout": 3000,
		"url":     "somethirdparty.com",
		"key":     12342134,
	}

	for key, value := range config {
		fmt.Printf("Config %s : %v\n", key, value)
	}
}

/*

Config timeout : 3000
Config url : somethirdparty.com
Config key : 12342134

*/
