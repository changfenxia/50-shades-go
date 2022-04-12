package main

import "fmt"

// #16. Strings are immutable

func main() {
	x := "teemo"

	// Error: cannot assign to x[0]
	// x[0] = 'T'

	// right version
	// it works in this case, but not always correct, some symbols will take more than 1 byte
	b := []byte(x)
	b[0] = 'T'

	// prints "Teemo"
	fmt.Println(string(b))
}