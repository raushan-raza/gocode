package main

import (
	"fmt"
	"testmod/fact"
)

func main() {

	var n, factorial int

	fmt.Println("enter n to calculate factorial")

	fmt.Scan(&n)

	factorial = fact.Factorial(n)

	fmt.Printf("fcatorial of n is %d\n", factorial)
}
