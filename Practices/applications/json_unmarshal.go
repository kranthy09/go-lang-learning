package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     string
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	fmt.Println(s)

	// conversion to go data structure
	// create slice of type person
	// person struct is defined by the keys in the string we are going to convert.
	var people []person

	// Unmarshal takes type slice of bytes, so string to byte conversion is made
	err := json.Unmarshal([]byte(s), &people)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(people) // prints the converted json in go data structure
}
