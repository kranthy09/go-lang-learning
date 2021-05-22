package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("before sort", xi)
	// sort xi
	sort.Ints(xi) // method to sort slice of int
	fmt.Println("after sort", xi)

	fmt.Println("before sort", xs)
	// sort xs
	sort.Strings(xs) // method to sort slice of string
	fmt.Println("after sort", xs)
}
