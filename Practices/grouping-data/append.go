package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// append an element to slice
	x = append(x, 56)
	fmt.Println(x)

	// append more than one element to slice
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{78, 43, 23, 89, 100}

	// append slice to another slice
	x = append(x, y...)
	fmt.Println(x)
}
