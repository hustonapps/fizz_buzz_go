package main

import "math/rand"

func main() {
	x := rand.Intn(45) + 1
	println(x)
	if x % 3 == 0 {
		if x % 5 == 0 {
			println("FizzBuzz")
		} else {
			println("Fizz")
		}
	} else if x % 5 == 0 {
		println("Buzz")
	}
}