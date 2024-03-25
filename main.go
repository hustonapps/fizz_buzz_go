package main

import (
	"fmt"
	"math/rand"
)

func getNum() int {
	x := rand.Intn(45) + 1
	return x
}

func printFizz() {
	fmt.Print("Fizz")
}

func printBuzz() {
	fmt.Print("Buzz")
}

func main() {
	for i := 0; i < 51; i++ {
		x := getNum()
		fmt.Printf("%v \n", x)
		if x % 3 == 0 {
			printFizz()
		}
		if x % 5 == 0 {
			printBuzz()
		}
		println()
		println()
	}
}