package main

import "fmt"

func main() {
	// create a SLICE using composite literal
	x := []int{3, 5, 2, 6, 7, 8, 9, 10}

	fmt.Println(x[:3])  // prints the values up to index 2
	fmt.Println(x[4:])  // prints the values from index to 4 to last element in the slice
	fmt.Println(x[1:5]) // prints values from index 1 to 4
	fmt.Println(x[:])   // prints complete slice
}
