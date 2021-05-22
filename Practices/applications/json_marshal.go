package main

import (
	"encoding/json"
	"fmt"
)

// create a user struct
type user struct {
	First string
	Age   int
}

func main() {
	// create users, u1, u2, u3 of type user
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	// create slice of user's by composite literal
	users := []user{u1, u2, u3}

	fmt.Println(users)

	// conversion to json string
	bs, err := json.Marshal(users) // returns byte string of json and error (if any)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(bs)) // prints json
}
