package main

import "fmt"

func main() {
	// create a map, with key as TYPE string and value as TYPE slice STRING
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// print out the records
	for key := range x {
		fmt.Println("For record: ", key)
		for index, value := range x[key] {
			fmt.Println("index: ", index, "value: ", value)
		}
	}

	// add record to map
	x["k_k"] = []string{`nobita`, "doreman", "skywalker"}

	// print the added record
	fmt.Println("Added record: ", x["k_k"])

	// delete the record
	delete(x, "no_dr")

	// print the map x
	fmt.Println(x)
}
