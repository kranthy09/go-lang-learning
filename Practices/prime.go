package main

import (
	"fmt"
	"math"
)

func main() {
	for n := 3; n  < 1000; n += 2{
		prime := true

		sqrtPlus1 := math.Sqrt(float64(n))
		for i := 3; float64(i) < sqrtPlus1 + 1; i += 2 {
			if n%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			fmt.Println(n)
		}
	}
}
