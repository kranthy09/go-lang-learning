package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 4, 3, 67, 109}

	// delete from the slice
	x = x[2:] // drop first two values from the slice
	fmt.Println(x)

	x = append(x[:3], x[6:]...) // drop values from index 3 to 6
	fmt.Println(x)

	y := x[:5]
	y = append(y, x[7])
	fmt.Println(y)
}
