package main

import "fmt"

func main() {
	s := []int{0,1,2,3,4,5,6,7,8,9,10}
	
	l := s[2:3]
	fmt.Println(l)

	l = s[:3]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)
}
