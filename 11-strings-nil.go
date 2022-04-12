package main

import "fmt"

// #11. Strings Can't Be "nil"

func main() {
	// wrong version	
	// Error: cannot use nil as type string in assignment
	// var s string = nil

	// Error: invalid operation: x == nil (mismatched types string and nil)
	// if s == nil {
	// 	s = "default"
	// }

	// right version
	var s string

	if s == "" {
		s = "default"
	}

	fmt.Print(s)
}