package main

// #14. Creating multi-dimensional structures

import "fmt"

func main() {
	h, w := 3,6

	raw := make([]int, h*w)
	for i := range raw {
		raw[i] = i
	}
	
	table := make([][]int, h)
	for i := range table {
		table[i] = raw[i*w:i*w+w]
	}

	// prints [[0 1 2 3 4 5] [6 7 8 9 10 11] [12 13 14 15 16 17]]
	fmt.Println(table)
}