package main

import "fmt"

func main() {
	xMap := map[string]int{
		"Bill": 39,
		"Tammy": 37,
	}

	if newUser, ok := xMap["James"]; ok {
		fmt.Printf("James is %d years old\n", newUser)
	} else {
		fmt.Printf("James not found\n")
	}
}