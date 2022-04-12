package main

// #12. Array Function Arguments (how to pass by reference)

import "fmt"

func main() {
	arr := [2]string{"Batman", "Robin"}

	// prints [Batman Robin], orinal array wasn't changed,
	// because by default it is passed by value
  func(x [2]string) {
		x[1] = "Joker"
	}(arr)
	fmt.Println(arr)
	
	// prints [Batman Joker] because we used a pointer to an original array
	func(x *[2]string) {
		(*x)[1] = "Joker"
	}(&arr)
	fmt.Println(arr)

	// we can also use slices to change the original array
	func(x []string) {
		x[1] = "Batwoman"
	}(arr[:])
	fmt.Println(arr)
}