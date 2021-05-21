package main

import "fmt"

// create a function, when we run the func, product will be prepared by given ingredients
func prepareProduct(ingredients string) string {
	return "Ingredients mixed: " + ingredients
}

// create a function which takes prepareProduct as argument
func manufacture(prepareProduct func(ingredients string) string) {
	fmt.Println(prepareProduct("Clove, aloe vera, ginger"))
	fmt.Println("Proceded to next step")
	fmt.Println("Product is manfactured and sealed.")
}

func main() {
	// call manufacture func to prepare and seal the product for sale.
	manufacture(prepareProduct)
}
