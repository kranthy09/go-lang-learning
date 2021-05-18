package main

import "fmt"

func main() {
	// create a slice using composite literal
	x := make([]string, 50, 500)
	for i := 0; i < len(x); i++ {
		x[i] = fmt.Sprintf("Position %v", i)
	}

	fmt.Println("slice x: ", x)
	fmt.Println("length of x:", len(x))
	fmt.Println("capacity of x:", cap(x))

	// append the values to slice, grows the length of the slice
	// the underlying capacity of slice 500 remains same
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`)
	fmt.Println("**slice properties after appending**")
	fmt.Println("slice x: ", x)
	fmt.Println("length of x:", len(x))
	fmt.Println("capacity of x:", cap(x))

	// prints the values, along with index
	for i := 0; i < len(x); i++ {
		fmt.Printf("index: %v value at index: %v\n", i, x[i])
	}
}
