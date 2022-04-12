package main

import "fmt"

// #6. Can't Use Short Variable Declarations to Set Field Values

type info struct {
	status bool
}

func getData() (bool, error) {
	return true,nil
}
func main() {
	var data info

	// doesn't work
	// data.result, err := getData()

	// so we have to use a temporary variable err
	var err error
	data.status, err = getData()

	if err == nil {
		fmt.Println("Info:", data.status)
	}
}