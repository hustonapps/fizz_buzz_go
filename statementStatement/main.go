package main

import (
	"fmt"
	"math/rand"
)

func main() {
	testVal := 33
	if val := rand.Intn(55); val > testVal {
		fmt.Printf("Value is greater than %v\n", testVal)
	} else if val == testVal {
		fmt.Printf("Value is %v\n", testVal)
	} else {
		fmt.Printf("Value is less than %v\n", testVal)
	}
}