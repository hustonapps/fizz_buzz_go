package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(45) + 1
	println(x)
	if x % 3 == 0 {
		fmt.Print("Fizz")
	}
	if x % 5 == 0 {
		fmt.Print("Buzz")
	}
	println()
}