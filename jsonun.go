package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
    Name string  `json:"name"`
    Email string `json:"email"`
    Age int64    `json:"age"`
}

func main() {
	var u User
	b := `{"name":"Alice","email":"alice@email.net","age":21}`
	err := json.Unmarshal([]byte(b), &u)
	if err != nil {
		panic(err)
	}

	fmt.Println("User", u)
}
