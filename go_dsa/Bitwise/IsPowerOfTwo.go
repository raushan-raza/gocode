package main

import (
	Bitpackage "Bitwise/MyBitPackages"
	"fmt"
)

func main() {

	var N int

	fmt.Println("Enter N to check if given no is power of 2 or not")
	fmt.Scan(&N)

	fmt.Println(Bitpackage.IsPowerOfTwo(N))

}
