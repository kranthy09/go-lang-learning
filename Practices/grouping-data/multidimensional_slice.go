package main

import "fmt"

func main() {
	// create multi-dimensional slice using COMPOSITE LITERAL
	names := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(names)

	// range over the records in multi-dimensional slice
	for index, value := range names {
		fmt.Println("Record: ")
		fmt.Println(index, value)
		// range over each record
		fmt.Println("Each value in record: ")
		for index, value := range names[index] {
			fmt.Println(index, value)
		}
	}
}
