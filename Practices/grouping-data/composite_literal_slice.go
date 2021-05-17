//1. Using a COMPOSITE LITERAL:
//		a. create a SLICE of TYPE int
//		b. assign 10 VALUES
// 2. Range over the slice and print the values out.
// 3. Using format printing
//		a. print out the TYPE of the slice

package main

import (
	"fmt"
)

func main() {
	// create a slice using COMPOSITE LITETAL
	numbers := []int{1, 3, 5, 2, 6, 7, 8, 313, 44, 5}
	fmt.Println(numbers)

	// range over the slice
	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}

	// prints out TYPE of slice
	fmt.Printf("TYPE of numbers: %T\n", numbers)
}
