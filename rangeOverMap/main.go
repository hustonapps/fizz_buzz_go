package main

import "fmt"

func main() {
	xMap := map[string]int{
		"Bill": 39,
		"Tammy": 37,
	}

	for k, v := range xMap {
		fmt.Printf("%s is %d years old\n", k, v)
	}
}