package main

import "fmt"

// #15. Checking for Non-Existing Map Keys

func main() {

	// wrong version. Prints "No entry"
	m := map[string]string{"one": "teemo", "two": "", "three": "nunu"}

	if v := m["two"]; v == "" {
		fmt.Println("No entry")
	}

	// right version. Prints "Entry found"
	// 
	if _, ok := m["four"]; !ok {
		fmt.Println("No entry")
	} else {
		fmt.Println("Entry found")
	}
}