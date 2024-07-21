package main

import (
	Bitpackage "Bitwise/MyBitPackages"
	"fmt"
)

func main() {

	var n int

	fmt.Println("Enter N to count set bits")
	fmt.Scan(&n)

	println(Bitpackage.CountBits(n))
}
