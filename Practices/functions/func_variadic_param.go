package main

import (
	"fmt"
)

// create a function which takes variadic parameter
func sum_of_numbers(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6}

	// passing slice x value to func by unfrilling it with "..."
	result := sum_of_numbers(x...)
	fmt.Println(result) // prints the result that returned when the func is called
}
