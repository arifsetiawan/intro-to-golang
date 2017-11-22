package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
    Name string  `json:"name"`
    Email string `json:"email"`
    Age int64    `json:"age,omitempty"`
}

func main() {
	//u := User{"Alice", "alice@email.net", 21}
	u := User{Name:"Alice", Email:"alice@email.net"}
	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println("User", string(b))
}
