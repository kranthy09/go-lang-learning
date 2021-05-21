package main

import "fmt"

func main() {
	// create an anounymous function
	func(x int) {
		for i := 0; i < x; i++ {
			fmt.Println(i)
		}
	}(10) // calling the function with 25 as parameter
}
